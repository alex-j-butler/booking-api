package rpc

import (
	"github.com/Qixalite/booking-api/config"
	"github.com/Qixalite/booking-api/server"
)

type V1ListServersReq struct {
	Key string
	Tag string
}

type V1ListServersResp struct {
	Servers          []V1Server
	ServersAvailable int
}

type V1Server struct {
	UUID       string
	Name       string
	Tags       []string
	Executable string
	Options    server.OptionKeyValues

	IPAddress string
	Port      int
	STVPort   int

	ServerPassword string
	RCONPassword   string

	Running bool
}
