package lnet

import (
	"net/http"

	"github.com/cheetah-fun-gs/goso/internal/biz/handler"
	"github.com/cheetah-fun-gs/goso/internal/common"
	"github.com/cheetah-fun-gs/goso/internal/generated"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/gin-gonic/gin"
)

// handleErrorSoNet 框架层错误处理
func handleErrorSoNet(code int, err error) interface{} {
	switch code {
	case http.StatusBadRequest:
		return handler.GetCommonResp(handler.CommonRespCodeClientUnknown, err)
	default:
		return handler.GetCommonResp(handler.CommonRespCodeServerUnknown, err)
	}
}

// SoHTTP 获取 gnet http 服务
func SoHTTP() (*sohttp.SoHTTP, error) {
	s, err := sohttp.NewLNet()
	if err != nil {
		return nil, err
	}

	// 注册首页
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to goso")
	})

	if err := s.SetConfig(&sohttp.Config{Ports: common.PortsHTTPLNet}); err != nil {
		return nil, err
	}

	if err := s.SetErrorNetFunc(handleErrorSoNet); err != nil {
		return nil, err
	}

	// 最后注册 handler
	for _, h := range generated.Handlers {
		if h.IsAnyNet(so.NetTypeLNet) {
			s.Register(h)
		}
	}
	return s, nil
}
