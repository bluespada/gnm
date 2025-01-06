package node

import (
	"fmt"
	"slices"

	"github.com/bluespada/gnm/internal/downloader"
	"github.com/bluespada/gnm/internal/filesystem"
	"github.com/bytedance/sonic"
)

var dist = "https://nodejs.org/dist/index.json"

var downloaded = "https://nodejs.org/dist/[version]/node-[version]-[platform]-[arch].tar.gz"

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

