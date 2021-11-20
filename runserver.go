package main

import (
	"files-server/conf"
	"files-server/routers"

	"github.com/julienschmidt/httprouter"
)

func Start(port string, dirpath string) *httprouter.Router {

	conf.Port = port
	conf.DirPath = dirpath

	return routers.Router()

}
