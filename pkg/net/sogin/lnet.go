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
		return err
	}

	err = json.Unmarshal(rawPack, req)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return err
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

func lnetConverFunc(handler so.Handler) gin.HandlerFunc {
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

// NewLnetGin 一个新的lnet gin 对象
func NewLnetGin(ports []int) (*SoGin, error) {
	lnet, err := New(ports)
	if err != nil {
		return nil, err
	}
	err = lnet.SetConverFunc(lnetConverFunc)
	if err != nil {
		return nil, err
	}
	return lnet, nil
}
