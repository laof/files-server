package models

import (
	"github.com/laof/goport"

	"github.com/laof/filesserver/conf"
)

func GetHostAddress() string {
	return "http://" + goport.GetIP() + ":" + conf.Port
}
