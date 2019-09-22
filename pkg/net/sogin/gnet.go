package sogin

import (
	"context"
	"encoding/json"
	"fmt"
	"goso/pkg/so"
	"goso/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gnetBeforeHandleFunc
func gnetParseRequest(c *gin.Context, req interface{}) error {
	rawPack, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil
	}

	err = json.Unmarshal(rawPack, req)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil
	}

	return nil
}

// gnetGetContextFunc 默认的获取 ctx 的方法
func gnetGetContextFunc(c *gin.Context) (context.Context, error) {
	data := map[string]interface{}{}
	for _, p := range c.Params {
		if strings.HasPrefix(p.Key, ContextPrefix) {
			data[strings.Replace(p.Key, ContextPrefix, "", 1)] = p.Value
		}
	}
	return utils.LoadContext(data), nil
}

func gnetGetHandlerFunc(handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				c.Status(http.StatusBadGateway) // 不建议使用 http code, 这是一个demo
				return
			}
		}()
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()

		ctx, err := gnetGetContextFunc(c)
		if err != nil {
			return
		}
		err = gnetParseRequest(c, req)
		if err != nil {
			return
		}
		err = handler.Func()(ctx, req, resp)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// GnetGin gin 版 gnet
type GnetGin struct {
	*SoGin
	GetHandlerFunc func(handle so.Handler) gin.HandlerFunc // so.HandlerFunc to gin.HandlerFunc
}

// Register 注册处理器
func (gnet *GnetGin) Register(handler so.Handler) error {
	privateData := handler.GetPrivateData().(*HandlerPrivateData)
	httpMethod := privateData.HTTPMethod
	uri := privateData.URI

	gnet.Handle(httpMethod, uri, gnet.GetHandlerFunc(handler))
	return nil
}

// NewgnetGin 一个新的gnet gin 对象
func NewgnetGin(ports []int) (*GnetGin, error) {
	gnet, err := New(ports)
	if err != nil {
		return nil, err
	}
	return &GnetGin{
		SoGin:          gnet,
		GetHandlerFunc: gnetGetHandlerFunc,
	}, nil
}
