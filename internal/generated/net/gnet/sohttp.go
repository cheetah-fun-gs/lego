package gnet

import (
	"context"
	"goso/internal/biz/handlers"
	"goso/pkg/handler"
	sohttp "goso/pkg/net/sohttp"
)

// handler 定义
var (
	CommonTimeSoGinHandler = &handler.Handler{
		Name:    "common.time",
		Routers: sohttp.NewRouters(handlers.CommonTimeURIS, handlers.CommonTimeHTTPMethods),
		Req:     &handlers.CommonTimeReq{},
		Resp:    &handlers.CommonTimeResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers.CommonTimeHandle(ctx, req.(*handlers.CommonTimeReq), resp.(*handlers.CommonTimeResp))
		},
	}
)

// SoHTTP 获取 gnet http 服务
func SoHTTP() (*sohttp.SoHTTP, error) {
	s, err := sohttp.NewGnet([]int{8080})
	if err != nil {
		return nil, err
	}
	s.Register(CommonTimeSoGinHandler)
	return s, nil
}
