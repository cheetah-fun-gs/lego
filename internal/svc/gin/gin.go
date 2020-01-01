package gin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/cheetah-fun-gs/lego/internal/biz/handler"
	allhandler "github.com/cheetah-fun-gs/lego/internal/generated/handler"
	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
	legogin "github.com/cheetah-fun-gs/lego/pkg/svc/gin"
	"github.com/gin-gonic/gin"
)

func isMatchString(s string, ss []string) bool {
	if len(ss) == 0 {
		return true
	}
	for _, a := range ss {
		if s == a {
			return true
		}
	}
	return false
}

func beforeHandleFunc(ctx context.Context, c *gin.Context, req interface{}) error {
	reqData, err := c.GetRawData()
	if err != nil {
		return err
	}

	if len(reqData) == 0 {
		return fmt.Errorf("rawData is blank")
	}

	if err := json.Unmarshal(reqData, req); err != nil {
		return err
	}

	// 获取公共请求头
	commonReq := reflect.ValueOf(req).Elem().FieldByName("Common").Interface().(*handler.CommonReq)
	// 合并客户端request id
	requestids := []string{}
	if commonReq.RequestID != "" {
		requestids = append(requestids, commonReq.RequestID)
	}
	if c.GetString(legogin.LegoRequestID) != "" {
		requestids = append(requestids, c.GetString(legogin.LegoRequestID))
	}
	commonReq.RequestID = strings.Join(requestids, "-")
	c.Set(legogin.LegoRequestID, strings.Join(requestids, "-"))
	return nil
}

func behindHandleFunc(ctx context.Context, c *gin.Context, resp interface{}) error {
	// 获取公共响应头
	commonResp := reflect.ValueOf(resp).Elem().FieldByName("Common").Interface().(*handler.CommonResp)
	// 填充requestID
	commonResp.RequestID = c.GetString(legogin.LegoRequestID)
	// 设置返回code和msg 用于日志中间件
	c.Set(logRespCode, commonResp.Code)
	c.Set(logRespMsg, commonResp.Msg)

	respData, _ := json.Marshal(resp)
	c.Data(http.StatusOK, "application/json; charset=utf-8", respData)
	return nil
}

// Start 启动一个gin服务
func Start(name string, addrs ...string) {
	engine := gin.New()

	// 应用中间件
	engine.Use(gin.Recovery(), middlewareCORS, middlewareLogger, middlewareLegoGin)

	// 注册首页
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to legogin")
	})

	// 最后注册 handler
	hs := []legocore.Handler{}
	for _, h := range allhandler.Handlers {
		svcNames := reflect.ValueOf(h).Elem().FieldByName("SvcNames").Interface().([]string)
		if isMatchString(name, svcNames) {
			hs = append(hs, h)
		}
	}

	legogin.Register(engine, beforeHandleFunc, behindHandleFunc, hs...)

	ginStart(engine, addrs...)
}
