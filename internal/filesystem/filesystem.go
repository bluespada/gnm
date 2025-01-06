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
    ExistOrCreate(nodeDist)
    node, _ := os.ReadDir(nodeDist)
    for _, e := range node {
        installed = append(installed, e.Name())
    }
    return installed
}

func ExistOrCreate(path string) {
    println(path)
    if _, err := os.Stat(path); os.IsNotExist(err) {
        os.MkdirAll(path, 0755)
    }
}

func GetDistFolder() string {
    return filepath.Join(root, nodeDist)
}

func init(){
    root, _ = os.UserHomeDir()
}
