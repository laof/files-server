package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/laof/fs/api"
	"github.com/laof/fs/conf"
	"github.com/laof/fs/models"
	"github.com/laof/goport"
)

//go:embed assets/index.html
var index_html string

//go:embed assets/ico.d6jo8a5r9k.svg
var ico_svg string

//go:embed assets/main.771b0e1f5c7ff2f5.js
var main_js string

//go:embed assets/polyfills.4ceba69a5f6f91c3.js
var polyfills_js string

//go:embed assets/runtime.1a1bb1bd7fe02d90.js
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
	r := api.Router()

	go (func() {
		time.Sleep(time.Second * 1)
		openbrowser("http://localhost:" + strconv.Itoa(port))
	})()

	e := http.ListenAndServe(":"+conf.Port, r)

	if e != nil {
		fmt.Println("server fail")
	}
}

func init() {
	models.Fmap = map[string]string{
		"/index.html":                    index_html,
		"/ico.d6jo8a5r9k.svg":            ico_svg,
		"/main.771b0e1f5c7ff2f5.js":      main_js,
		"/polyfills.4ceba69a5f6f91c3.js": polyfills_js,
		"/runtime.1a1bb1bd7fe02d90.js":   runtime_js,
		"/styles.a241b4a8abcb89ef.css":   styles_css,
	}
}

func openbrowser(url string) {

	// go func() {
	// 	recover()
	// }()

	var err error
	switch runtime.GOOS {
	case "linux":
		// err = exec.Command("xdg-open", url).Start()
		fmt.Println("linux ....")
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
