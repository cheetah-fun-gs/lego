package sogin

import (
	"context"
	"encoding/json"
	"goso/pkg/so"
	"goso/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// DefaultHandler 默认的sogin的处理器
type DefaultHandler struct {
	Name       string
	URI        string
	HTTPMethod string
	Req        interface{} // 请求结构体指针
	Resp       interface{} // 响应结构体指针
	Func       func(ctx context.Context, req, resp interface{}) error
}

// GetName 获取处理器名称
func (h *DefaultHandler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *DefaultHandler) GetRouter() string {
	return h.URI
}

// GetReq 获取请求结构体
func (h *DefaultHandler) GetReq() interface{} {
	return h.Req
}

// GetResp 获取响应结构体
func (h *DefaultHandler) GetResp() interface{} {
	return h.Resp
}

// Handle 处理器方法
func (h *DefaultHandler) Handle(ctx context.Context, req, resp interface{}) error {
	return h.Func(ctx, req, resp)
}

// HandlerPrivateData 获取私有数据
type HandlerPrivateData struct {
	HTTPMethod string
}

// GetPrivateData 获取私有数据
func (h *DefaultHandler) GetPrivateData() interface{} {
	return &HandlerPrivateData{
		HTTPMethod: h.HTTPMethod,
	}
}

// New 默认sogin对象
func New() (*SoGin, error) {
	router := gin.Default()
	return &SoGin{
		Engine: router,
		NetAttr: &so.NetAttr{
			Ports: []int{8000},
		},
		GetContextFunc:   DefaultGetContextFunc,
		BeforeHandleFunc: DefaultBeforeHandleFunc,
		AfterHandleFunc:  DefaultAfterHandleFunc,
	}, nil
}

// DefaultBeforeHandleFunc 默认的预处理
func DefaultBeforeHandleFunc(c *gin.Context, req interface{}) error {
	rawPack, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return nil
	}

	err = json.Unmarshal(rawPack, req)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return nil
	}

	return nil
}

// DefaultAfterHandleFunc 默认的结束处理
func DefaultAfterHandleFunc(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

// DefaultGetContextFunc 默认的获取 ctx 的方法
func DefaultGetContextFunc(c *gin.Context) (context.Context, error) {
	data := map[string]interface{}{}
	for _, p := range c.Params {
		if strings.HasPrefix(p.Key, ContextPrefix) {
			data[strings.Replace(p.Key, ContextPrefix, "", 1)] = p.Value
		}
	}
	return utils.LoadContext(data), nil
}
