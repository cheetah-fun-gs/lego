// Package sogin 符合goso net对象的 gin 服务
package sogin

import (
	"context"
	"fmt"
	"goso/pkg/so"

	"github.com/gin-gonic/gin"
)

// HandlerPrivateData 获取私有数据
type HandlerPrivateData struct {
	HTTPMethod string
}

// SoGin 符合goso net对象的 gin 服务
type SoGin struct {
	Ports []int
	*gin.Engine
	GetContextFunc   func(c *gin.Context) (context.Context, error)
	BeforeHandleFunc func(c *gin.Context, req interface{}) error
	AfterHandleFunc  func(c *gin.Context, resp interface{}) error
}

// SetBeforeHandleFunc 设置before handle
func (soGin *SoGin) SetBeforeHandleFunc(beforeHandleFunc func(c *gin.Context, req interface{}) error) {
	soGin.BeforeHandleFunc = beforeHandleFunc
}

// SetAfterHandleFunc 设置after handle
func (soGin *SoGin) SetAfterHandleFunc(afterHandleFunc func(c *gin.Context, resp interface{}) error) {
	soGin.AfterHandleFunc = afterHandleFunc
}

// SetGetContextFunc 设置get context func
func (soGin *SoGin) SetGetContextFunc(getContextFunc func(c *gin.Context) (context.Context, error)) {
	soGin.GetContextFunc = getContextFunc
}

// Register 注册处理器
func (soGin *SoGin) Register(handler so.Handler) {
	privateData := handler.GetPrivateData().(*HandlerPrivateData)
	httpMethod := privateData.HTTPMethod

	req := handler.GetReq()
	resp := handler.GetResp()

	soGin.Handle(
		httpMethod,
		handler.GetRouter(),
		func(c *gin.Context) {
			ctx, err := soGin.GetContextFunc(c)
			if err != nil {
				return
			}
			err = soGin.BeforeHandleFunc(c, req)
			if err != nil {
				return
			}
			err = handler.Handle(ctx, req, resp)
			if err != nil {
				return
			}
			err = soGin.AfterHandleFunc(c, resp)
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
	for _, port := range soGin.Ports {
		addr = append(addr, fmt.Sprintf(":%d", port))
	}
	return soGin.Run(addr...)
}

// Stop 关闭服务
func (soGin *SoGin) Stop() error {
	return soGin.Stop()
}
