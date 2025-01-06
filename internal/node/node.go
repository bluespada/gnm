package node

import (
    "io"
    "log"
    "net/http"
    "os"
    "fmt"

    "github.com/schollz/progressbar/v3"
)

var dist = "https://nodejs.org/dist/index.json"

var downloaded = "https://nodejs.org/dist/[version]/node-[version]-[platform]-[arch].tar.gz"


type nodeVersionDownload struct {
    version string
    platform string
    arch string
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


func Fetch(ltsOnly bool) {
    req, err := http.NewRequest("GET", dist, nil)
    if err != nil {
        log.Fatalln("Failed to fetch node version:", err)
    }

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()

    f, _ := os.OpenFile("node_dist.json", os.O_WRONLY|os.O_CREATE, 0644)
    defer f.Close()

    bar := progressbar.DefaultBytes(
        res.ContentLength,
        "downloading",
    )

    io.Copy(io.MultiWriter(f, bar), res.Body)
}

func DownloadAndInstall(version string) {
    req, err := http.NewRequest("GET", dist, nil)
    if err != nil {
        log.Fatalln("Failed to fetch node version:", err)
    }

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()

    f, _ := os.OpenFile("node_dist.json", os.O_WRONLY|os.O_CREATE, 0644)
    defer f.Close()

    bar := progressbar.DefaultBytes(
        res.ContentLength,
    )

    io.Copy(io.MultiWriter(f, bar), res.Body)
}
