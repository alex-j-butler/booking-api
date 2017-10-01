package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1StopServerReq struct {
	Key  string
	UUID string
}

func init() {
	gorpc.RegisterType(&V1StopServerReq{})
}
