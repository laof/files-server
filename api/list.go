package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilesData struct {
	Children  []FilesData `json:"children"`
	Download  string      `json:"download"`
	Extension string      `json:"extension"`
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Type      string      `json:"type"`
	Size      int64       `json:"size"`
}

func ListDir(dirPth string, fs *FilesData) (err error) {

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {

		path := filepath.ToSlash(dirPth + PthSep + fi.Name())

		f := FilesData{
			[]FilesData{},
			path[len(rootFile):],
			filepath.Ext(fi.Name()),
			fi.Name(),
			path,
			fi.Name(),
			fi.Size()}

		if fi.IsDir() {
			f.Type = "directory"
			s, _ := DirSize(f.Path)
			f.Size = s
			ListDir(f.Path, &f)
			fs.Children = append(fs.Children, f)

		} else {
			f.Type = "file"
			f.Children = nil
			fs.Children = append(fs.Children, f)
		}
	}
	return nil
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func getPath() string {

	str, _ := os.Getwd()
	file := filepath.ToSlash(str)

	return file
}

var rootFile string

func init() {
	// file := getPath()
	// http.FileServer(http.Dir(file))

	rootFile = getPath()
}

type JsonData struct {
	FilesData
	Success bool `json:"success"`
}

func GetList() []byte {

	path := rootFile

	f := FilesData{
		[]FilesData{},
		"",
		"",
		path,
		path,
		"directory",
		0}

	s, _ := DirSize(f.Path)

	f.Size = s

	_ = ListDir(f.Path, &f)

	v, _ := json.Marshal(struct {
		FilesData
		Success bool `json:"success"`
	}{
		FilesData(f),
		true,
	})

	return v
}
