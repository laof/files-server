package routers

import (
	"github.com/laof/filesserver/conf"

	"github.com/julienschmidt/httprouter"
)

func Start(port string, dirpath string) *httprouter.Router {

	conf.Port = port
	conf.DirPath = dirpath

	return Router()
}
