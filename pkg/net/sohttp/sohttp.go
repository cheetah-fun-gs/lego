// Package sohttp 符合goso net对象的 http 服务
package sohttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/logger"
	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

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
	if soHTTP.ErrorNetFunc == nil {
		// 未定义 http code 的处理回调, 直接使用 http 错误码, 不建议
		c.Status(code)
		return
	}

	// http code 的处理回调
	c.JSON(http.StatusOK, soHTTP.ErrorNetFunc(code, err))
	return
}

// Config 配置
type Config struct {
	Ports []int
}

// SoHTTP 符合goso net对象的 http 服务
// PreHandleFunc handle 的 前置处理
// PostHandleFunc handle 的 后置处理
// PreHandleFunc 和 PostHandleFunc 返回的 code 由 ErrorNetFunc 处理
// 不设置 ErrorNetFunc, code 只能使用 httpcode, 默认使用 BadRequest 和 InternalServerError
type SoHTTP struct {
	*gin.Engine
	Config           *Config
	Logger           so.Logger
	GatePack         so.GatePack // gnet 用到
	ErrorNetFunc     func(code int, err error) interface{}
	PreHandleFunc    func(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, int, error)
	PostHandleFunc   func(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) (int, error)
	ConverHandleFunc func(soHTTP *SoHTTP, handle so.Handler) gin.HandlerFunc // so.HandlerFunc to gin.HandlerFunc
}

// SetLogger 设置日志器
func (soHTTP *SoHTTP) SetLogger(logger so.Logger) {
	soHTTP.Logger = logger
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

// SetPreHandleFunc 设置 响应编码方法
func (soHTTP *SoHTTP) SetPreHandleFunc(preHandleFunc func(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, int, error)) error {
	soHTTP.PreHandleFunc = preHandleFunc
	return nil
}

// SetPostHandleFunc 设置 请求解码方法
func (soHTTP *SoHTTP) SetPostHandleFunc(postHandleFunc func(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) (int, error)) error {
	soHTTP.PostHandleFunc = postHandleFunc
	return nil
}

// SetConverHandleFunc 设置 ConverHandleFunc
func (soHTTP *SoHTTP) SetConverHandleFunc(ConverHandleFunc func(soHTTP *SoHTTP, handle so.Handler) gin.HandlerFunc) error {
	soHTTP.ConverHandleFunc = ConverHandleFunc
	return nil
}

// SetErrorNetFunc 设置框架层错误处理
func (soHTTP *SoHTTP) SetErrorNetFunc(errFunc func(code int, err error) interface{}) error {
	soHTTP.ErrorNetFunc = errFunc
	return nil
}

// Register 注册处理器
func (soHTTP *SoHTTP) Register(handler so.Handler) error {
	routers := handler.GetRouter()
	for _, router := range routers {
		r := router.(*Router)
		soHTTP.Handle(r.HTTPMethod, r.URI, soHTTP.ConverHandleFunc(soHTTP, handler))
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

func defaultPreHandleFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, int, error) {
	ctx := context.Background()

	rawPack, err := c.GetRawData()
	if err != nil {
		soHTTP.Logger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return ctx, http.StatusBadRequest, err
	}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, req)
		if err != nil {
			soHTTP.Logger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return ctx, http.StatusBadRequest, err
		}
	}

	ctx = context.WithValue(ctx, ContextKey("trace_id"), fmt.Sprintf("%v", uuid.NewV4()))
	return ctx, http.StatusOK, nil
}

func defaultPostHandleFunc(ctx context.Context, soHTTP *SoHTTP, c *gin.Context, resp interface{}) (int, error) {
	c.JSON(http.StatusOK, resp)
	return http.StatusOK, nil
}

func defaultConverHandleFunc(soHTTP *SoHTTP, handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, ContextKey("uri"), c.Request.URL.Path)

		defer func() {
			if r := recover(); r != nil {
				soHTTP.Logger.Error(ctx, "InternalServerError defaultConverHandleFunc error: %v", r)
				errorHandle(ctx, soHTTP, c, http.StatusInternalServerError, fmt.Errorf("%v", r))
				return
			}
		}()

		ctx, code, err := soHTTP.PreHandleFunc(soHTTP, c, req)
		if err != nil {
			soHTTP.Logger.Error(ctx, "BadRequest defaultPreHandleFunc error: %v", err)
			errorHandle(ctx, soHTTP, c, code, err)
			return
		}

		if err := handler.Handle(ctx, req, resp); err != nil {
			soHTTP.Logger.Error(ctx, "InternalServerError %v Handle error: %v", handler.GetName(), err)
			errorHandle(ctx, soHTTP, c, http.StatusInternalServerError, err)
			return
		}

		if code, err := soHTTP.PostHandleFunc(ctx, soHTTP, c, resp); err != nil {
			soHTTP.Logger.Error(ctx, "InternalServerError defaultPostHandleFunc error: %v", err)
			errorHandle(ctx, soHTTP, c, code, err)
			return
		}
		return
	}
}

// New 默认http对象
func New() (*SoHTTP, error) {
	router := gin.Default()
	soHTTP := &SoHTTP{
		Engine:           router,
		Config:           &Config{},
		Logger:           &logger.Logger{},
		PreHandleFunc:    defaultPreHandleFunc,
		PostHandleFunc:   defaultPostHandleFunc,
		ConverHandleFunc: defaultConverHandleFunc,
	}
	return soHTTP, nil
}
