package models

import (
	"github.com/laof/goport"

	"github.com/laof/fs/conf"
)

func GetHostAddress() string {
	ip := ""
	if ips := goport.GetIP(); len(ips) > 0 {
		ip = ips[0]
	}
	return "http://" + ip + ":" + conf.Port
}
