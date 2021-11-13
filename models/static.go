package models

import (
	"github.com/gobuffalo/packr/v2"
)

func GetStaticMapping() (smap map[string]string) {
	smap = make(map[string]string)
	box := packr.New("static", "../static")

	stc := []string{"main.html", "main.ffc5d1007650b2b6b303.js", "polyfills.80e47a8e2355f1ceab2c.js", "runtime.91a02d1be9f8a4fe75ea.js", "styles.47822e33bf6c3a9b00bc.css"}

	for _, filename := range stc {
		txt, _ := box.FindString(filename)
		smap["/"+filename] = txt
	}

	return
}
