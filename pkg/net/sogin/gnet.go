package sogin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GateHandler gate的sogin的处理器
type GateHandler struct {
	Name       string
	URI        string
	HTTPMethod string
	Req        interface{} // 请求结构体指针
	Resp       interface{} // 响应结构体指针
	Func       func(ctx context.Context, req, resp interface{}) error
}

// GetName 获取处理器名称
func (h *GateHandler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *GateHandler) GetRouter() string {
	return h.URI
}

// GetReq 获取请求结构体
func (h *GateHandler) GetReq() interface{} {
	return h.Req
}

// GetResp 获取响应结构体
func (h *GateHandler) GetResp() interface{} {
	return h.Resp
}

// Handle 处理器方法
func (h *GateHandler) Handle(ctx context.Context, req, resp interface{}) error {
	return h.Func(ctx, req, resp)
}

// GetPrivateData 获取私有数据
func (h *GateHandler) GetPrivateData() interface{} {
	return &HandlerPrivateData{
		HTTPMethod: h.HTTPMethod,
	}
}

// NewGate 一个新的gate sogin对象
func NewGate(ports []int) (*SoGin, error) {
	soGin, err := New(ports)
	if err != nil {
		return nil, err
	}
	return soGin, nil
}

// GateBeforeHandleFunc gate 预处理
func GateBeforeHandleFunc(c *gin.Context, req interface{}) error {
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

// GateAfterHandleFunc gate 结束处理
func GateAfterHandleFunc(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

// GateGetContextFunc gate 不从外部获取 context, 自己生产
func GateGetContextFunc(c *gin.Context) (context.Context, error) {
	ctx := context.Background() // :TODO
	return ctx, nil
}
