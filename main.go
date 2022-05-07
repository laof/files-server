package main

import (
	"fmt"
	"net/http"

	"github.com/laof/filesserver/conf"
	"github.com/laof/filesserver/routers"
	"github.com/laof/goport"
)

func main() {

	conf.Port = goport.InputPort(conf.Port)
	r := routers.Router()

	e := http.ListenAndServe(":"+conf.Port, r)

	if e != nil {
		fmt.Println("server fail")
	}
	fmt.Println("end")
}
