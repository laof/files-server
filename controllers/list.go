package controllers

import (
	"files-server/libs"
	"files-server/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type filesInfo struct {
	*models.Files
	Success bool `json:"success"`
}

func GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	files := models.GetDirInfo()
	w.Write(libs.JsonData(filesInfo{files, true}))
}
