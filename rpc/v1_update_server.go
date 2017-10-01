package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1UpdateServerReq struct {
	Key  string
	UUID string
}

func init() {
	gorpc.RegisterType(&V1UpdateServerReq{})
}
