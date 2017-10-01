package rpc

import (
	"github.com/Qixalite/booking-api/config"
	"github.com/Qixalite/booking-api/server"
	"github.com/valyala/gorpc"
)

type V1NextServerReq struct {
	Key string
	Tag string
}

type V1NextServerResp struct {
	V1Server
}

func init() {
	gorpc.RegisterType(&V1NextServerReq{})
	gorpc.RegisterType(&V1NextServerResp{})
}
