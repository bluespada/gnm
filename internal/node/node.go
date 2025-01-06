package node


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
