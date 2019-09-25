package gnet

import (
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/so"

	"github.com/cheetah-fun-gs/goso/internal/biz/handlers"
	"github.com/cheetah-fun-gs/goso/internal/common"
	"github.com/cheetah-fun-gs/goso/internal/generated"

	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
	"github.com/gin-gonic/gin"
)

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

	if err := s.SetConfig(&sohttp.Config{Ports: common.PortsHTTPGNet}); err != nil {
		return nil, err
	}

	if err := s.SetErrorNetFunc(handlers.HandleCommonRespSoNet); err != nil {
		return nil, err
	}

	// 最后注册 handler
	for _, h := range generated.Handlers {
		if h.IsAnyNet(so.NetTypeGNet) {
			s.Register(h)
		}
	}
	return s, nil
}
