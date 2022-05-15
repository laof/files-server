package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/laof/filesserver/conf"
	"github.com/laof/filesserver/models"
	"github.com/laof/filesserver/routers"
	"github.com/laof/goport"
)

//go:embed assets/main.html
var main_html string

//go:embed assets/main.cfb326e733b33e067096.js
var main_cfb326e733b33e067096_js string

//go:embed assets/polyfills.4fcb186b1abbc00c95e4.js
var polyfills_4fcb186b1abbc00c95e4_js string

//
//go:embed assets/runtime.91a02d1be9f8a4fe75ea.js
var runtime_91a02d1be9f8a4fe75ea_js string

//go:embed assets/styles.47822e33bf6c3a9b00bc.css
var styles_47822e33bf6c3a9b00bc_css string

func main() {
	assets()
	conf.Port = goport.InputPort(conf.Port)
	r := routers.Router()

	e := http.ListenAndServe(":"+conf.Port, r)

	if e != nil {
		fmt.Println("server fail")
	}
}

func assets() {
	models.Fmap = map[string]string{
		"/main.html":                         main_html,
		"/main.cfb326e733b33e067096.js":      main_cfb326e733b33e067096_js,
		"/polyfills.4fcb186b1abbc00c95e4.js": polyfills_4fcb186b1abbc00c95e4_js,
		"/runtime.91a02d1be9f8a4fe75ea.js":   runtime_91a02d1be9f8a4fe75ea_js,
		"/styles.47822e33bf6c3a9b00bc.css":   styles_47822e33bf6c3a9b00bc_css,
	}
}
