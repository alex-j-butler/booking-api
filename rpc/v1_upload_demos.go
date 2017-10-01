package rpc

import (
	"errors"

	"github.com/Qixalite/booking-api/config"
	"github.com/valyala/gorpc"
)

type V1UploadDemosReq struct {
	Key      string
	UUID     string
	Uploader string
}

type V1UploadDemosResp struct {
	Demos []string
}

func init() {
	gorpc.RegisterType(&V1UploadDemosReq{})
	gorpc.RegisterType(&V1UploadDemosResp{})
}
