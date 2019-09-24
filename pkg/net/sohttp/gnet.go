package sohttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cheetah-fun-gs/goso/pkg/gatepack"
	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/cheetah-fun-gs/goso/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gnetBeforeHandleFunc
func gnetParseRequest(c *gin.Context, req interface{}) (ctx context.Context, err error) {
	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest GetRawData error: %v", err)
		return nil, err
	}

	gatePack := &gatepack.JSONPack{}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, gatePack)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			return nil, err
		}
	}

	err = gatePack.Verify()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest Verify error: %v", err)
		return nil, err
	}

	if len(gatePack.LogicPack) != 0 {
		err = json.Unmarshal(gatePack.LogicPack, req)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
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

func gnetConverFunc(config *Config, handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				soLogger.Error(ctx, "BadGateway gnetConverFunc error: %v", r)
				errorHandle(c, config, http.StatusBadGateway, fmt.Errorf("%v", r))
				return
			}
		}()

		ctx, err := gnetParseRequest(c, req)
		if err != nil {
			soLogger.Error(ctx, "BadGateway gnetParseRequest error: %v", err)
			errorHandle(c, config, http.StatusBadRequest, err)
			return
		}

		err = handler.Handle(ctx, req, resp)
		if err != nil {
			soLogger.Error(ctx, "BadGateway %v Handle error: %v", handler.GetName(), err)
			errorHandle(c, config, http.StatusBadGateway, err)
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// NewGnet 一个新的gnet gin 对象
func NewGnet() (*SoHTTP, error) {
	gnet, err := New()
	if err != nil {
		return nil, err
	}
	err = gnet.SetConverFunc(gnetConverFunc)
	if err != nil {
		return nil, err
	}
	return gnet, nil
}
