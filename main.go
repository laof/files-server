package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/laof/filesserver/conf"
	"github.com/laof/filesserver/models"
	"github.com/laof/filesserver/routers"
	"github.com/laof/goport"
)

//go:embed assets/index.tw3d2elq9cw2pq6t025.html
var index_tw3d2elq9cw2pq6t025_html string

//go:embed assets/main.cfb326e73c3b3e3e067096.js
var main_cfb326e73c3b3e3e067096_js string

//go:embed assets/polyfills.4fcb186b1abbc00c95e4.js
var polyfills_4fcb186b1abbc00c95e4_js string

//go:embed assets/runtime.91a02d1be9f8a4fe75ea.js
var runtime_91a02d1be9f8a4fe75ea_js string

//go:embed assets/styles.47822e33bf6c3a9b00bc.css
var styles_47822e33bf6c3a9b00bc_css string

func main() {

	port := 6200
	args := os.Args

	if len(args) > 1 {
		inp := args[1]
		n, e := strconv.Atoi(inp)
		if e == nil {
			if n >= 2000 && n <= 49152 {
				port = n
			} else {
				fmt.Print("input error:  >= 2000 && <= 49152")
				return
			}
		}
	}

	assets()
	conf.Port = goport.InputPort(strconv.Itoa(port))
	r := routers.Router()

	e := http.ListenAndServe(":"+conf.Port, r)

	if e != nil {
		fmt.Println("server fail")
	}
}

func assets() {
	models.Fmap = map[string]string{
		"/index.tw3d2elq9cw2pq6t025.html":    index_tw3d2elq9cw2pq6t025_html,
		"/main.cfb326e73c3b3e3e067096.js":    main_cfb326e73c3b3e3e067096_js,
		"/polyfills.4fcb186b1abbc00c95e4.js": polyfills_4fcb186b1abbc00c95e4_js,
		"/runtime.91a02d1be9f8a4fe75ea.js":   runtime_91a02d1be9f8a4fe75ea_js,
		"/styles.47822e33bf6c3a9b00bc.css":   styles_47822e33bf6c3a9b00bc_css,
	}
}
