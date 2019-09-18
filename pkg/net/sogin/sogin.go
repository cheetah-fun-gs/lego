// Package sogin 符合goso net对象的 gin 服务
package sogin

import (
	"context"
	"fmt"
	"goso/pkg/so"

	"github.com/gin-gonic/gin"
)

// SoGin 符合goso net对象的 gin 服务
type SoGin struct {
	*gin.Engine
	*so.NetAttr
	BeforeHandleFunc func(c *gin.Context) (ctx context.Context, req interface{}, err error)
	AfterHandleFunc  func(resp interface{}) error
}

// SetAttr 设置属性
func (soGin *SoGin) SetAttr(attr *so.NetAttr) {
	soGin.NetAttr = attr
}

// SetBeforeHandleFunc 设置before handle
func (soGin *SoGin) SetBeforeHandleFunc(beforeHandleFunc func(c *gin.Context) (ctx context.Context, req interface{}, err error)) {
	soGin.BeforeHandleFunc = beforeHandleFunc
}

// SetAfterHandleFunc 设置after handle
func (soGin *SoGin) SetAfterHandleFunc(afterHandleFunc func(resp interface{}) error) {
	soGin.AfterHandleFunc = afterHandleFunc
}

// Register 注册处理器
func (soGin *SoGin) Register(handler *Handler) {
	soGin.Handle(
		handler.HTTPMethod,
		handler.URI,
		func(c *gin.Context) {
			ctx, req, err := soGin.BeforeHandleFunc(c)
			if err != nil {
				return
			}
			resp, err := handler.Handle(ctx, req)
			if err != nil {
				return
			}
			err = soGin.AfterHandleFunc(resp)
			if err != nil {
				return
			}
			return
		},
	)
}

// Start 启动服务
func (soGin *SoGin) Start() error {
	addr := []string{}
	for _, port := range soGin.NetAttr.Ports {
		addr = append(addr, fmt.Sprintf(":%d", port))
	}
	return soGin.Run(addr...)
}

// Stop 关闭服务
func (soGin *SoGin) Stop() error {
	return soGin.Stop()
}
