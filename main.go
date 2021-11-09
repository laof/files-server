package main

import (
	"files-server/api"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/laof/goport"
)

type HttpServe struct {
}

func (h HttpServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, cke := r.Cookie("fs")

	if cke != nil {
		ck := &http.Cookie{
			Name:     "fs",
			Value:    strconv.FormatInt(time.Now().UnixNano(), 10),
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			MaxAge:   0}
		http.SetCookie(w, ck)
	}

	path := r.URL.Path
	m := strings.ToLower(r.Method)

	if path == "/" {
		w.Write([]byte(staticMap["/index.html"]))
	} else if m == "post" {
		if path == "/api/host" {
			w.Write(api.GetHost())
		} else if path == "/api/list" {
			w.Write(api.GetList())
		} else if path == "/api/upload" {
			w.Write(api.UploadFiles(w, r))
		} else if path == "/api/talk_history" {
			w.Write(api.GetTakHistory())
		}

	} else {

		if path == "/index.html" {
			http.ServeFile(w, r, filepath.Join("./", path))
		} else if staticMap[path] != "" {

			if strings.Contains(path, ".css") {
				w.Header().Add("content-type", "text/css; charset=utf-8")
			} else {
				w.Header().Add("content-type", "application/javascript; charset=utf-8")
			}

			w.Write([]byte(staticMap[path]))
		}

	}

}

func main() {

	initStaticMap()

	socketServer := api.WebsocketServer()

	go socketServer.Serve()
	defer socketServer.Close()

	port := goport.InputPort("9873")
	api.SetPort(port)

	hs := new(HttpServe)

	http.Handle("/socket.io/", socketServer)
	http.Handle("/", hs)

	e := http.ListenAndServe(":"+port, nil)

	if e != nil {
		fmt.Println("server fail")
	}
}

var staticMap = map[string]string{}

func initStaticMap() map[string]string {
	box := packr.New("static", "./static")

	statics := []string{"index.html",
		"main.ffc5d1007650b2b6b303.js",
		"polyfills.80e47a8e2355f1ceab2c.js",
		"runtime.91a02d1be9f8a4fe75ea.js",
		"styles.47822e33bf6c3a9b00bc.css"}

	for _, filename := range statics {
		txt, _ := box.FindString(filename)
		staticMap["/"+filename] = txt
	}

	return staticMap

}
