package downloader

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/schollz/progressbar/v3"
)


func Fetch(url string) (string, error) {

    req, _ := http.NewRequest("GET", url, nil)
    res, err := http.DefaultClient.Do(req)

    if err != nil {
        log.Fatalln(err)
        return "", err
    }

    defer res.Body.Close()


    strBody, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
        return "", err
    }

    return string(strBody), nil
}

func Download(url string, version string, path string) {

    req, _ := http.NewRequest("GET", url, nil)
    res, err := http.DefaultClient.Do(req)

    if err != nil {
        log.Fatalln(err)
        os.Exit(-1)
    }

    defer res.Body.Close()

    f, _ := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0644)
    defer f.Close()

    bar := progressbar.DefaultBytes(
        res.ContentLength,
        fmt.Sprintf("Download %s", version),
    )

    io.Copy(io.MultiWriter(f, bar), res.Body)

}
