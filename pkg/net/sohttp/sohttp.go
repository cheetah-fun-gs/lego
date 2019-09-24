// Package sohttp 符合goso net对象的 http 服务
package sohttp

import (
	"fmt"
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/logger"
	"github.com/cheetah-fun-gs/goso/pkg/so"

	"github.com/gin-gonic/gin"
)

var soLogger = logger.New()

func errorHandle(c *gin.Context, config *Config, code int, err error) {
	if config.HTTPCodeFunc == nil {
		// 未定义 http code 的处理回调, 直接使用 http 错误码, 不建议
		c.Status(code)
		return
	}

	// http code 的处理回调
	c.JSON(http.StatusOK, config.HTTPCodeFunc(code, err))
	return
}

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
type ConverFunc func(config *Config, handle so.Handler) gin.HandlerFunc

// Config 配置
type Config struct {
	Ports        []int
	HTTPCodeFunc func(code int, err error) interface{} // 对 http 错误码的处理, BadRequest 和 BadGateway
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

// SetConfig 设置 Config
func (soHTTP *SoHTTP) SetConfig(config *Config) error {
	soHTTP.Config = config
	return nil
}

// Register 注册处理器
func (soHTTP *SoHTTP) Register(handler so.Handler) error {
	routers := handler.GetRouter()
	for _, router := range routers {
		r := router.(*Router)
		soHTTP.Handle(r.HTTPMethod, r.URI, soHTTP.ConverFunc(soHTTP.Config, handler))
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
func New() (*SoHTTP, error) {
	router := gin.Default()
	return &SoHTTP{
		Engine: router,
		Config: &Config{},
	}, nil
}
