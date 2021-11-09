package api

import (
	"encoding/json"
	"fmt"
	"net"
)

type HostData struct {
	Host     string `json:"host"`
	Shutdown bool   `json:"shutdown"`
}

var port string

func SetPort(st string) {
	port = st
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var ip string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return "http://" + ip + ":" + port
}

func GetHost() []byte {
	data := HostData{GetIP(), false}

	v, e := json.Marshal(data)

	if e != nil {
		return []byte("")
	}

	return v
}
