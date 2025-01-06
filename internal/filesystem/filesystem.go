package filesystem

import (
    "os"
    "path/filepath"
)

var root string
var nodeDist = ".gnm/dist/node"

func GetListInstalled() []string {
    var nodeDist = filepath.Join(root, nodeDist)
    var installed = []string{}
    existOrCreate(nodeDist)
    node, _ := os.ReadDir(nodeDist)
    for _, e := range node {
        installed = append(installed, e.Name())
    }
    return installed
}

func existOrCreate(path string) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        os.Mkdir(path, 0755)
    }
}


func init(){
    root, _ = os.UserHomeDir()
}
