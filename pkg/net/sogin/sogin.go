// Package sogin 符合goso net对象的 gin 服务
package sogin

import (
	"context"
	"fmt"
	"goso/pkg/so"

	"github.com/gin-gonic/gin"
)

// Handler sogin的处理器
type Handler struct {
	Name       string
	URI        string
	HTTPMethod string
	Req        interface{} // 请求结构体指针
	Resp       interface{} // 响应结构体指针
	Func       func(ctx context.Context, req, resp interface{}) error
}

// GetName 获取处理器名称
func (h *Handler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *Handler) GetRouter() string {
	return h.URI
}

// GetReq 获取请求结构体
func (h *Handler) GetReq() interface{} {
	return h.Req
}

// GetResp 获取响应结构体
func (h *Handler) GetResp() interface{} {
	return h.Resp
}

// Handle 处理器方法
func (h *Handler) Handle(ctx context.Context, req, resp interface{}) error {
	return h.Func(ctx, req, resp)
}

// GetPrivateData 获取私有数据
func (h *Handler) GetPrivateData() interface{} {
	return &HandlerPrivateData{
		HTTPMethod: h.HTTPMethod,
	}
}

// HandlerPrivateData 获取私有数据
type HandlerPrivateData struct {
	HTTPMethod string
	URI        string
}

// Config 配置
type Config struct {
	Ports []int
}

// SoGin 符合goso net对象的 gin 服务
type SoGin struct {
	*gin.Engine
	Config *Config
}

// Register 注册处理器
func (soGin *SoGin) Register(handler so.Handler) error {
	return fmt.Errorf("Not implemented")
}

// Start 启动服务
func (soGin *SoGin) Start() error {
	addr := []string{}
	for _, port := range soGin.Config.Ports {
		addr = append(addr, fmt.Sprintf(":%d", port))
	}
	return soGin.Run(addr...)
}

// Stop 关闭服务
func (soGin *SoGin) Stop() error {
	return soGin.Stop()
}

// New 默认sogin对象
func New(ports []int) (*SoGin, error) {
	router := gin.Default()
	return &SoGin{
		Engine: router,
		Config: &Config{
			Ports: ports,
		},
	}, nil
}
