package conf

import (
	"os"
	"path/filepath"
)

var (
	Port    = "9873"
	DirPath = ""
)

func init() {
	str, _ := os.Getwd()
	DirPath = filepath.ToSlash(str)
}
