// Package sogin 符合goso net对象的 gin 服务
package sogin

import (
	"fmt"
	"goso/pkg/so"

	"github.com/gin-gonic/gin"
)

// Router 路由器
type Router struct {
	HTTPMethod string
	URI        string
}

// String 格式化方法
func (router *Router) String() string {
	return fmt.Sprintf("%v-%v", router.HTTPMethod, router.HTTPMethod)
}

// NewRouters 获取路由
func NewRouters(uris []string, httpMethods []string) []interface{} {
	routers := []interface{}{}
	for _, httpMethod := range httpMethods {
		for _, uri := range uris {
			routers = append(routers, &Router{
				URI:        uri,
				HTTPMethod: httpMethod,
			})
		}
	}
	return routers
}

// ConverFunc so.HandlerFunc to gin.HandlerFunc
type ConverFunc func(handle so.Handler) gin.HandlerFunc

// Config 配置
type Config struct {
	Ports []int
}

// SoGin 符合goso net对象的 gin 服务
type SoGin struct {
	*gin.Engine
	Config     *Config
	ConverFunc ConverFunc
}

// SetConverFunc 设置 ConverFunc
func (soGin *SoGin) SetConverFunc(converFunc ConverFunc) error {
	soGin.ConverFunc = converFunc
	return nil
}

// Register 注册处理器
func (soGin *SoGin) Register(handler so.Handler) error {
	routers := handler.GetRouter()
	for _, router := range routers {
		r := router.(*Router)
		soGin.Handle(r.HTTPMethod, r.URI, soGin.ConverFunc(handler))
	}
	return nil
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
