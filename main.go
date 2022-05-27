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

//go:embed assets/index.html
var index_html string

//go:embed assets/ico.d6jo8a5r9k.svg
var ico_svg string

//go:embed assets/main.7b900f1b2c601e51.js
var main_js string

//go:embed assets/polyfills.1b42a168c772535f.js
var polyfills_js string

//go:embed assets/runtime.3989356f74c4e649.js
var runtime_js string

//go:embed assets/styles.a241b4a8abcb89ef.css
var styles_css string

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

	conf.Port = goport.InputPort(strconv.Itoa(port))
	r := routers.Router()

	e := http.ListenAndServe(":"+conf.Port, r)

	if e != nil {
		fmt.Println("server fail")
	}
}

func init() {
	models.Fmap = map[string]string{
		"/index.html":                    index_html,
		"/ico.d6jo8a5r9k.svg":            ico_svg,
		"/main.7b900f1b2c601e51.js":      main_js,
		"/polyfills.1b42a168c772535f.js": polyfills_js,
		"/runtime.3989356f74c4e649.js":   runtime_js,
		"/styles.a241b4a8abcb89ef.css":   styles_css,
	}
}
