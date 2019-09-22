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

// lnetBeforeHandleFunc
func lnetParseRequest(c *gin.Context, req interface{}) error {
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

// lnetGetContextFunc 默认的获取 ctx 的方法
func lnetGetContextFunc(c *gin.Context) (context.Context, error) {
	data := map[string]interface{}{}
	for key, val := range c.Request.PostForm {
		if strings.HasPrefix(key, ContextPrefix) {
			if len(val) == 1 {
				data[strings.Replace(key, ContextPrefix, "", 1)] = val[0]
			}
		}
	}
	return utils.LoadContext(data), nil
}

func lnetGetHandlerFunc(handler so.Handler) gin.HandlerFunc {
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

		ctx, err := lnetGetContextFunc(c)
		if err != nil {
			return
		}
		err = lnetParseRequest(c, req)
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

// LnetGin gin 版 lnet
type LnetGin struct {
	*SoGin
	GetHandlerFunc func(handle so.Handler) gin.HandlerFunc // so.HandlerFunc to gin.HandlerFunc
}

// Register 注册处理器
func (lnet *LnetGin) Register(handler so.Handler) error {
	privateData := handler.GetPrivateData().(*HandlerPrivateData)
	httpMethod := privateData.HTTPMethod
	uri := privateData.URI

	lnet.Handle(httpMethod, uri, lnet.GetHandlerFunc(handler))
	return nil
}

// NewLnetGin 一个新的lnet gin 对象
func NewLnetGin(ports []int) (*LnetGin, error) {
	lnet, err := New(ports)
	if err != nil {
		return nil, err
	}
	return &LnetGin{
		SoGin:          lnet,
		GetHandlerFunc: lnetGetHandlerFunc,
	}, nil
}
