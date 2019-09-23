package sohttp

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
func gnetParseRequest(c *gin.Context, req interface{}) (ctx context.Context, err error) {
	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest GetRawData error: %v", err)
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil, err
	}

	gatePack := &gatepack.JSONPack{}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, gatePack)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
			return nil, err
		}
	}

	err = gatePack.Verify()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest Verify error: %v", err)
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
		return nil, err
	}

	if len(gatePack.LogicPack) != 0 {
		err = json.Unmarshal(gatePack.LogicPack, req)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			c.Status(http.StatusBadRequest) // 不建议使用 http code, 这是一个demo
			return nil, err
		}
	}

	ctx = gatePack.GetContext()
	return ctx, nil
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
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				soLogger.Error(ctx, "BadGateway gnetConverFunc error: %v", r)
				c.Status(http.StatusBadGateway) // 不建议使用 http code, 这是一个demo
				return
			}
		}()

		ctx, err := gnetParseRequest(c, req)
		if err != nil {
			soLogger.Error(ctx, "BadGateway gnetParseRequest error: %v", err)
			c.Status(http.StatusBadGateway) // 不建议使用 http code, 这是一个demo
			return
		}
		fmt.Println(req)
		err = handler.Handle(ctx, req, resp)
		if err != nil {
			soLogger.Error(ctx, "BadGateway %v Handle error: %v", handler.GetName(), err)
			c.Status(http.StatusBadGateway) // 不建议使用 http code, 这是一个demo
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// NewGnet 一个新的gnet gin 对象
func NewGnet(ports []int) (*SoHTTP, error) {
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
