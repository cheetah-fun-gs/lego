// Package sohttp 符合goso net对象的 http 服务
package sohttp

import (
	"fmt"
	"goso/pkg/logger"
	"goso/pkg/so"

	"github.com/gin-gonic/gin"
)

var soLogger = logger.New()

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

// SoHTTP 符合goso net对象的 http 服务
type SoHTTP struct {
	*gin.Engine
	Config     *Config
	ConverFunc ConverFunc
}

// SetConverFunc 设置 ConverFunc
func (soHTTP *SoHTTP) SetConverFunc(converFunc ConverFunc) error {
	soHTTP.ConverFunc = converFunc
	return nil
}

// Register 注册处理器
func (soHTTP *SoHTTP) Register(handler so.Handler) error {
	routers := handler.GetRouter()
	for _, router := range routers {
		r := router.(*Router)
		soHTTP.Handle(r.HTTPMethod, r.URI, soHTTP.ConverFunc(handler))
	}
	return nil
}

// Start 启动服务
func (soHTTP *SoHTTP) Start() error {
	addr := []string{}
	for _, port := range soHTTP.Config.Ports {
		addr = append(addr, fmt.Sprintf(":%d", port))
	}
	return soHTTP.Run(addr...)
}

// Stop 关闭服务
func (soHTTP *SoHTTP) Stop() error {
	return soHTTP.Stop()
}

// New 默认http对象
func New(ports []int) (*SoHTTP, error) {
	router := gin.Default()
	return &SoHTTP{
		Engine: router,
		Config: &Config{
			Ports: ports,
		},
	}, nil
}
