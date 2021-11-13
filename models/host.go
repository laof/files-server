package models

import (
	"github.com/laof/goport"

	"files-server/conf"
)

func GetHostAddress() string {
	return "http://" + goport.GetIP() + ":" + conf.Port
}
