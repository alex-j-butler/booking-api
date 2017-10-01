package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1TagReq struct {
	Key  string
	UUID string
	Tag  string
}

type V1TagResp struct {
	Success bool
}

func init() {
	gorpc.RegisterType(&V1TagReq{})
	gorpc.RegisterType(&V1TagResp{})
}
