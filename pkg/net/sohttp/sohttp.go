// Package sohttp 符合goso net对象的 http 服务
package sohttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/logger"
	"github.com/cheetah-fun-gs/goso/pkg/so"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

var soLogger = logger.New()

// ContextKey ctx key
type ContextKey interface{}

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

func errorHandle(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, code int, err error) {
	if soHTTP.HTTPCodeFunc == nil {
		// 未定义 http code 的处理回调, 直接使用 http 错误码, 不建议
		c.Status(code)
		return
	}

	// http code 的处理回调
	c.JSON(http.StatusOK, soHTTP.HTTPCodeFunc(ctx, soHTTP, code, err))
	return
}

// Config 配置
type Config struct {
	Ports []int
}

// SoHTTP 符合goso net对象的 http 服务
type SoHTTP struct {
	*gin.Engine
	Config        *Config
	GatePack      so.GatePack                                                                // gnet 用到
	HTTPCodeFunc  func(ctx context.Context, soHTTP *SoHTTP, code int, err error) interface{} // 对 http 错误码的处理, BadRequest 和 BadGateway
	UnmarshalFunc func(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, error)
	MarshalFunc   func(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) error
	ConverFunc    func(soHTTP *SoHTTP, handle so.Handler) gin.HandlerFunc // so.HandlerFunc to gin.HandlerFunc
}

// SetConfig 设置 Config
func (soHTTP *SoHTTP) SetConfig(config *Config) error {
	soHTTP.Config = config
	return nil
}

// SetGatePack 设置包对象
func (soHTTP *SoHTTP) SetGatePack(gatePack so.GatePack) error {
	soHTTP.GatePack = gatePack
	return nil
}

// SetHTTPCodeFunc 设置 HTTPCodeFunc
func (soHTTP *SoHTTP) SetHTTPCodeFunc(httpCodeFunc func(ctx context.Context, soHTTP *SoHTTP, code int, err error) interface{}) error {
	soHTTP.HTTPCodeFunc = httpCodeFunc
	return nil
}

// SetUnmarshalFunc 设置 响应编码方法
func (soHTTP *SoHTTP) SetUnmarshalFunc(unmarshalFunc func(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, error)) error {
	soHTTP.UnmarshalFunc = unmarshalFunc
	return nil
}

// SetMarshalFunc 设置 请求解码方法
func (soHTTP *SoHTTP) SetMarshalFunc(marshalFunc func(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) error) error {
	soHTTP.MarshalFunc = marshalFunc
	return nil
}

// SetConverFunc 设置 ConverFunc
func (soHTTP *SoHTTP) SetConverFunc(converFunc func(soHTTP *SoHTTP, handle so.Handler) gin.HandlerFunc) error {
	soHTTP.ConverFunc = converFunc
	return nil
}

// Register 注册处理器
func (soHTTP *SoHTTP) Register(handler so.Handler) error {
	routers := handler.GetRouter()
	for _, router := range routers {
		r := router.(*Router)
		soHTTP.Handle(r.HTTPMethod, r.URI, soHTTP.ConverFunc(soHTTP, handler))
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

// GetPrivateData 获取私有数据
func (soHTTP *SoHTTP) GetPrivateData() interface{} {
	return nil
}

func defaultUnmarshalFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, error) {
	ctx := context.Background()

	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return ctx, err
	}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, req)
		if err != nil {
			soLogger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return ctx, err
		}
	}

	context.WithValue(ctx, ContextKey("trace_id"), fmt.Sprintf("%v", uuid.NewV4()))
	return ctx, nil
}

func defaultMarshalFunc(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

func defaultConverFunc(soHTTP *SoHTTP, handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				soLogger.Error(ctx, "BadGateway defaultConverFunc error: %v", r)
				errorHandle(ctx, soHTTP, c, http.StatusBadGateway, fmt.Errorf("%v", r))
				return
			}
		}()

		ctx, err := soHTTP.UnmarshalFunc(soHTTP, c, req)
		if err != nil {
			soLogger.Error(ctx, "BadRequest defaultUnmarshalFunc error: %v", err)
			errorHandle(ctx, soHTTP, c, http.StatusBadRequest, err)
			return
		}

		if err := handler.Handle(ctx, req, resp); err != nil {
			soLogger.Error(ctx, "BadGateway %v Handle error: %v", handler.GetName(), err)
			errorHandle(ctx, soHTTP, c, http.StatusBadGateway, err)
			return
		}

		if err := soHTTP.MarshalFunc(ctx, soHTTP, c, resp); err != nil {
			soLogger.Error(ctx, "BadGateway defaultMarshalFunc error: %v", err)
			errorHandle(ctx, soHTTP, c, http.StatusBadGateway, err)
			return
		}
		return
	}
}

// New 默认http对象
func New() (*SoHTTP, error) {
	router := gin.Default()
	soHTTP := &SoHTTP{
		Engine:        router,
		Config:        &Config{},
		UnmarshalFunc: defaultUnmarshalFunc,
		MarshalFunc:   defaultMarshalFunc,
		ConverFunc:    defaultConverFunc,
	}
	return soHTTP, nil
}
