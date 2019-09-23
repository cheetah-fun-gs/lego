package gnet

import (
	"context"
	"goso/internal/biz/handlers"
	"goso/pkg/net/sogin"
)

// handler 定义
var (
	CommonTimeSoGinHandler = &sogin.Handler{
		Name:        "common.time",
		URIS:        handlers.CommonTimeURIS,
		HTTPMethods: handlers.CommonTimeURIS,
		Req:         &handlers.CommonTimeReq{},
		Resp:        &handlers.CommonTimeResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers.CommonTimeHandle(ctx, req.(*handlers.CommonTimeReq), resp.(*handlers.CommonTimeResp))
		},
	}
)

// SoGin 获取 gnet sogin 服务
func SoGin() (*sogin.SoGin, error) {
	s, err := sogin.NewGnet([]int{8080})
	if err != nil {
		return nil, err
	}
	s.Register(CommonTimeSoGinHandler)
	return s, nil
}
