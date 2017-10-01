package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1SendCommandReq struct {
	Key     string
	UUID    string
	Command string
}

func init() {
	gorpc.RegisterType(&V1SendCommandReq{})
}
