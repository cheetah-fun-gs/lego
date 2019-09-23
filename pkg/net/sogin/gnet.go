package sogin

import (
	"context"
	"encoding/json"
	"fmt"
	"goso/pkg/gatepack"
	"goso/pkg/so"
	"goso/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gnetBeforeHandleFunc
func gnetParseRequest(c *gin.Context) (ctx context.Context, req interface{}, err error) {
	rawPack, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil, nil, err
	}

	gatePack := &gatepack.JSONPack{}
	err = json.Unmarshal(rawPack, gatePack)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil, nil, err
	}

	err = gatePack.Verify()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil, nil, err
	}

	ctx = gatePack.GetContext()
	req = gatePack.GetLogicPack()
	return ctx, req, nil
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

func gnetConverFunc(handler so.Handler) gin.HandlerFunc {
	resp := handler.GetResp()

	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				c.Status(http.StatusBadGateway) // 不建议使用 http code, 这是一个demo
				fmt.Println(err)
				return
			}
		}()
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()

		ctx, req, err := gnetParseRequest(c)
		if err != nil {
			return
		}
		err = handler.Handle(ctx, req, resp)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// NewGnet 一个新的gnet gin 对象
func NewGnet(ports []int) (*SoGin, error) {
	gnet, err := New(ports)
	if err != nil {
		return nil, err
	}
	err = gnet.SetConverFunc(gnetConverFunc)
	if err != nil {
		return nil, err
	}
	return gnet, nil
}
