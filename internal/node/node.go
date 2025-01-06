package node

import (
    "fmt"
    "runtime"
    "slices"

    "github.com/bluespada/gnm/internal/downloader"
    "github.com/bluespada/gnm/internal/filesystem"
    "github.com/bytedance/sonic"
    "github.com/schollz/progressbar/v3"
)

var dist = "https://nodejs.org/dist/index.json"

var downloadUrl = "https://nodejs.org/dist/%s/node-%s-%s.tar.gz"

var downloadStep = map[string]func(p any){
    "[cyan][1/3][reset] Download...": func(version any){
        var arch = GetArchitecture()
        var installedPath = filesystem.GetDistFolder()
        var targetPath = fmt.Sprintf("%s/%s", installedPath, version)
        filesystem.ExistOrCreate(targetPath)
        downloader.Download(
            fmt.Sprintf(
                downloadUrl,
                version,
                version,
                arch,
            ),
            version.(string),
            fmt.Sprintf("%s/%s.tar.gz", targetPath, version),
        )
    },
    "[cyan][2/3][reset] Extracting...": func(p any){

    },
    "[cyan][3/3][reset] Finishing...": func(p any){

    },
}

type nodeVersionInfo struct {
    Version  string   `json:"version"`
    Date     string   `json:"date"`
    Files    []string `json:"files"`
    Npm      string   `json:"npm"`
    V8       string   `json:"v8"`
    Uv       string   `json:"uv"`
    Zlib     string   `json:"zlib"`
    Openssl  string   `json:"openssl"`
    Modules  string   `json:"modules"`
    Lts      any     `json:"lts"`
    Security bool     `json:"security"`
}

func ListRemote() {
    raw_str, _ := downloader.Fetch(dist)
    installed := filesystem.GetListInstalled()
    var dist = []nodeVersionInfo{}
    sonic.UnmarshalString(raw_str, &dist)
    slices.Reverse(dist)

    for _, i := range dist {
        var isInstalledS string
        var isInstalledD string
        var isLts string
        for _, n := range installed {
            if i.Version == n {
                isInstalledS = "[\033[32m+\033[0m]"
                isInstalledD = "(\033[32mInstalled\033[0m)"
            }else{
                isInstalledS = "[\033[31m-\033[0m]"
                isInstalledD = "(\033[31mNot Installed\033[0m)"
            }
        }
        if i.Lts == false {
            isLts = ""
        }else{
            isLts = fmt.Sprintf("-> \033[36mlts/%s\033[0m", i.Lts)
        }
        fmt.Println(
            isInstalledS,
            i.Version,
            isLts,
            isInstalledD,
        )
    }
}

func ListInstalled() {
    installed := filesystem.GetListInstalled()
    for _, i := range installed {
        fmt.Println(
            "[-]",
            i,
            "(default)",
        )
    }
}

func DownloadAndInstall(version string) {
    bar := progressbar.NewOptions(
        len(downloadStep),
        progressbar.OptionEnableColorCodes(true),
        progressbar.OptionShowBytes(true),
        progressbar.OptionSetWidth(15),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "[green]=[reset]",
            SaucerHead:    "[green]>[reset]",
            SaucerPadding: " ",
            BarStart:      "[",
            BarEnd:        "]",
        }),
    )

    for n, i := range downloadStep {
        bar.Describe(n)
        bar.Add(1);
        i(version)
    }
}

func GetArchitecture() string {
    if runtime.GOOS == "linux" {
        var arch string
        if runtime.GOARCH == "amd64" {
            arch = "x64"
        }
        return fmt.Sprintf(
            "%s-%s",
            runtime.GOOS,
            arch,
        )
    }
    // no windows & mac support :v 
    return "Undefined"
}
