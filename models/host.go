package models

import (
	"github.com/laof/goport"

	"filesserver/conf"
)

func GetHostAddress() string {
	return "http://" + goport.GetIP() + ":" + conf.Port
}
