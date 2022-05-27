package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/laof/filesserver/conf"
	"github.com/laof/filesserver/models"
)

func IndexFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// http.ServeFile(w, r, filepath.Join("./", path))
	// http.ServeFile(w, r, filepath.Join(conf.DirPath, r.URL.Path))
}

type NotFoundHttpServe struct {
}

func (h *NotFoundHttpServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	sp := staticmap(path)

	if sp == "" {

		http.ServeFile(w, r, filepath.Join(conf.DirPath, path))

	} else {

		ct := "content-type"
		if strings.Contains(path, ".svg") {
			w.Header().Add(ct, "image/svg+xml")
		} else if strings.Contains(path, ".css") {
			w.Header().Add(ct, "text/css; charset=utf-8")
		} else {
			w.Header().Add(ct, "application/javascript; charset=utf-8")
		}

		w.Write([]byte(sp))

	}
}

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// _, err := r.Cookie("fs")
	// if err != nil {
	// 	ck := &http.Cookie{
	// 		Name:     "fs",
	// 		Value:    strconv.FormatInt(time.Now().UnixNano(), 10),
	// 		Path:     "/",
	// 		Secure:   true,
	// 		HttpOnly: true,
	// 		MaxAge:   0}
	// 	http.SetCookie(w, ck)
	// }
	w.Write([]byte(staticmap("/index.html")))
}

func staticmap(key string) string {
	return models.Fmap[key]
}
