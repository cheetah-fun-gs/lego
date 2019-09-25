package gnet

import (
	"context"
	"net/http"

	"github.com/cheetah-fun-gs/goso/internal/biz/handlers"
	handlerscommon "github.com/cheetah-fun-gs/goso/internal/biz/handlers/common"
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
	"github.com/gin-gonic/gin"
)

// handler 定义
var hs = []*handler.Handler{
	&handler.Handler{
		Name:    "CommonPing",
		Routers: sohttp.NewRouters(handlerscommon.PingURIS, handlerscommon.PingHTTPMethods),
		Req:     &handlerscommon.PingReq{},
		Resp:    &handlerscommon.PingResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlerscommon.PingHandle(ctx, req.(*handlerscommon.PingReq), resp.(*handlerscommon.PingResp))
		},
	},
	&handler.Handler{
		Name:    "CommonPost",
		Routers: sohttp.NewRouters(handlerscommon.PostURIS, handlerscommon.PostHTTPMethods),
		Req:     &handlerscommon.PostReq{},
		Resp:    &handlerscommon.PostResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlerscommon.PostHandle(ctx, req.(*handlerscommon.PostReq), resp.(*handlerscommon.PostResp))
		},
	},
}

// SoHTTP 获取 gnet http 服务
func SoHTTP() (*sohttp.SoHTTP, error) {
	s, err := sohttp.NewGNet()
	if err != nil {
		return nil, err
	}

	// 注册首页
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to goso")
	})

	if err := s.SetConfig(&sohttp.Config{Ports: []int{8080}}); err != nil {
		return nil, err
	}

	if err := s.SetErrorNetFunc(handlers.HandleCommonRespSoNet); err != nil {
		return nil, err
	}

	// 最后注册 handler
	for _, h := range hs {
		s.Register(h)
	}
	return s, nil
}
