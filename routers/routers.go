package routers

import (
	"github.com/laof/filesserver/controllers"

	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {

	r := httprouter.New()

	r.GET("/", controllers.Home)

	r.POST("/api/host", controllers.GetHost)
	r.POST("/api/list", controllers.GetList)
	r.POST("/api/upload", controllers.UploadFiles)
	r.POST("/api/talk_history", controllers.GetTakHistory)

	// r.GET("/index.html", controllers.IndexFiles)

	r.GET("/api/chat", controllers.ChatServer)

	r.NotFound = new(controllers.NotFoundHttpServe)

	return r

}
