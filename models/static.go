package models

import (
	"github.com/gobuffalo/packr/v2"
)

func GetStaticMapping() (smap map[string]string) {
	smap = make(map[string]string)
	box := packr.New("static", "../static")

	stc := []string{"main.html", "main.cfb326e733b33e067096.js", "polyfills.4fcb186b1abbc00c95e4.js", "runtime.91a02d1be9f8a4fe75ea.js", "styles.47822e33bf6c3a9b00bc.css"}

	for _, filename := range stc {
		txt, _ := box.FindString(filename)
		smap["/"+filename] = txt
	}

	return
}
