package controllers

import (
	"filesserver/models"
	"filesserver/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostData struct {
	Host     string `json:"host"`
	Shutdown bool   `json:"shutdown"`
}

func GetHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := HostData{models.GetHostAddress(), false}
	w.Write(utils.JsonData(data))
}
