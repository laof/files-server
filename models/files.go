package models

import (
	"github.com/laof/filesserver/conf"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Files struct {
	Children  []Files `json:"children"`
	Download  string  `json:"download"`
	Extension string  `json:"extension"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Type      string  `json:"type"`
	Size      int64   `json:"size"`
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func listDir(dirPth string, fs *Files) (err error) {

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {

		path := filepath.ToSlash(dirPth + PthSep + fi.Name())

		f := Files{
			[]Files{},
			path[len(conf.DirPath):],
			filepath.Ext(fi.Name()),
			fi.Name(),
			path,
			fi.Name(),
			fi.Size()}

		if fi.IsDir() {
			f.Type = "directory"
			s, _ := dirSize(f.Path)
			f.Size = s
			listDir(f.Path, &f)
			fs.Children = append(fs.Children, f)

		} else {
			f.Type = "file"
			f.Children = nil
			fs.Children = append(fs.Children, f)
		}
	}
	return nil
}

func GetDirInfo() *Files {
	path := conf.DirPath

	f := Files{
		[]Files{},
		"",
		"",
		path,
		path,
		"directory",
		0}

	s, _ := dirSize(f.Path)

	f.Size = s

	_ = listDir(f.Path, &f)

	return &f
}
