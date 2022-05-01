package main

import (
	"fmt"
	"net"
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

	fmt.Print(GetIP())
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String()
	}
	return ""

}
