package lnet

import (
	"context"
	"goso/internal/biz/handlers"
	"goso/pkg/handler"
	"goso/pkg/net/sogin"
)

// handler 定义
var (
	CommonTimeSoGinHandler = &handler.Handler{
		Name:    "common.time",
		Routers: sogin.NewRouters(handlers.CommonTimeURIS, handlers.CommonTimeHTTPMethods),
		Req:     &handlers.CommonTimeReq{},
		Resp:    &handlers.CommonTimeResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers.CommonTimeHandle(ctx, req.(*handlers.CommonTimeReq), resp.(*handlers.CommonTimeResp))
		},
	}
)

// SoGin 获取 gnet sogin 服务
func SoGin() (*sogin.SoGin, error) {
	s, err := sogin.NewGnet([]int{8000})
	if err != nil {
		return nil, err
	}
	s.Register(CommonTimeSoGinHandler)
	return s, nil
}
