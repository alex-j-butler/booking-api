package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1ListServerReq struct {
	Key  string
	UUID string
}

type V1ListServerResp struct {
	V1Server
}

func init() {
	gorpc.RegisterType(&V1ListServerReq{})
	gorpc.RegisterType(&V1ListServerResp{})
}
