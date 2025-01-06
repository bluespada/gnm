package downloader

import (
    "io"
    "log"
    "net/http"
    _ "github.com/schollz/progressbar/v3"
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
