package gin

import (
	"context"
	"fmt"
	"net/http"

	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
	"github.com/gin-gonic/gin"
)

// 常量
const (
	LegoHandlerErr = "lego-handler-err"
	LegoHandlerMsg = "lego-handler-msg"

	HandlerCrash       = "handler crash"
	HandlerError       = "handler error"
	BeforeHandlerError = "before handler error"
	BehindHandlerError = "behind handler error"
)

// Register 注册处理器
func Register(engine *gin.Engine, beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handlers ...legocore.Handler) {
	for _, h := range handlers {
		routers := h.GetRouter()
		for _, r := range routers {
			method := r.(*Router).Method
			path := r.(*Router).Path
			engine.Handle(method, path, converHandle(beforeHandle, behindHandle, h))
		}
	}
	return
}

func converHandle(beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handler legocore.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = ContextWithRouter(ctx, &Router{
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
		})

		defer func() {
			if r := recover(); r != nil {
				c.Set(LegoHandlerErr, HandlerCrash)
				c.Set(LegoHandlerMsg, fmt.Sprintf("%v", r))
				c.Status(http.StatusInternalServerError)
				return
			}
		}()

		req := handler.CloneReq()
		resp := handler.CloneResp()

		// 前置处理
		if beforeHandle != nil {
			if err := beforeHandle(ctx, c, req); err != nil {
				c.Set(LegoHandlerErr, BeforeHandlerError)
				c.Set(LegoHandlerMsg, err.Error())
				c.Status(http.StatusBadRequest)
				return
			}
		} else {
			req = c.Request
		}

		// 处理
		if err := handler.Handle(ctx, req, resp); err != nil {
			c.Set(LegoHandlerErr, HandlerError)
			c.Set(LegoHandlerMsg, err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		// 后置处理
		if behindHandle != nil {
			if err := behindHandle(ctx, c, resp); err != nil {
				c.Set(LegoHandlerErr, BehindHandlerError)
				c.Set(LegoHandlerMsg, err.Error())
				c.Status(http.StatusInternalServerError)
				return
			}
		} else {
			c.JSON(http.StatusOK, resp)
			return
		}
	}
}
