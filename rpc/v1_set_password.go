package rpc

import (
	"errors"
	"fmt"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1SetPasswordReq struct {
	Key            string
	UUID           string
	RCONPassword   string
	ServerPassword string
}

func init() {
	gorpc.RegisterType(&V1SetPasswordReq{})
}
