package sohttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/cheetah-fun-gs/goso/pkg/utils"
	"github.com/gin-gonic/gin"
)

// lnetBeforeHandleFunc
func lnetParseRequest(c *gin.Context, req interface{}) error {
	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest GetRawData error: %v", err)
		return err
	}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, req)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			return err
		}
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

func lnetConverFunc(soHTTP *SoHTTP, handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				soLogger.Error(ctx, "BadGateway lnetConverFunc error: %v", r)
				errorHandle(c, soHTTP, http.StatusBadGateway, fmt.Errorf("%v", r))
				return
			}
		}()

		ctx, err := lnetGetContextFunc(c)
		if err != nil {
			soLogger.Error(ctx, "BadGateway lnetGetContextFunc error: %v", err)
			errorHandle(c, soHTTP, http.StatusBadGateway, err)
			return
		}
		err = lnetParseRequest(c, req)
		if err != nil {
			soLogger.Error(ctx, "BadRequest lnetParseRequest error: %v", err)
			errorHandle(c, soHTTP, http.StatusBadRequest, err)
			return
		}

		err = handler.Handle(ctx, req, resp)

		if err != nil {
			soLogger.Error(ctx, "BadGateway %v Handle error: %v", handler.GetName(), err)
			errorHandle(c, soHTTP, http.StatusBadGateway, err)
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// NewLNet 一个新的lnet gin 对象
func NewLNet() (so.LNet, error) {
	lnet, err := New()
	if err != nil {
		return nil, err
	}

	if err := lnet.SetConverFunc(lnetConverFunc); err != nil {
		return nil, err
	}

	return lnet, nil
}
