package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Sizer interface {
	Size() int64
}

type Data struct {
	Success bool  `json:"success"`
	Size    int64 `json:"size"`
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func smartName(dirPath string, filename string, i int) string {

	path := filepath.Base(filename)
	ext := filepath.Ext(path)
	pre := path[0 : len(path)-len(ext)]

	if i == 0 {
		path = pre + ext
	} else {
		path = pre + strconv.Itoa(i) + ext
	}

	if pathExists(filepath.Join(dirPath, path)) {
		return smartName(dirPath, filename, i+1)
	} else {
		return path
	}
}

func init() {
	//p:=smartName("E:/go-app/github.com/laof/golang/test/vue.js进阶.rar",0)
	//fmt.Println("======>>>>",p)
}

func UploadFiles(w http.ResponseWriter, r *http.Request) []byte {

	//dir
	file, handler, err := r.FormFile("upload")
	dir := r.FormValue("dir")
	if err != nil {
		http.Error(w, err.Error(), 500)
		fail := Data{false, 0}
		f, _ := json.Marshal(fail)
		return f
	}

	defer file.Close()
	ff := smartName(dir, handler.Filename, 0)
	savePath := filepath.Join(dir, ff)
	fmt.Println("Upload => ", savePath)
	f, _ := os.Create(savePath)

	defer f.Close()
	io.Copy(f, file)

	data := Data{true, file.(Sizer).Size()}
	s, _ := json.Marshal(data)
	return s
}
