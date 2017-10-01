package rpc

import (
	"github.com/valyala/gorpc"
)

type V1ConsoleServerReq struct {
	Key   string
	UUID  string
	Lines int
}

type V1ConsoleServerResp struct {
	ConsoleLines []string
	Lines        int
}

func init() {
	gorpc.RegisterType(&V1ConsoleServerReq{})
	gorpc.RegisterType(&V1ConsoleServerResp{})
}
