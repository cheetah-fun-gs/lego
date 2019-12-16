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
	HandlerError = "lego-handler-error"
	HandlerMsg   = "lego-handler-msg"

	HandlerErrorCrash     = "crash"
	HandlerErrorInBefore  = "in before"
	HandlerErrorInBehind  = "in behind"
	HandlerErrorInProcess = "in process"
)

// Register 注册处理器
func Register(engine *gin.Engine, beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handlers ...legocore.Handler) {
	for _, h := range handlers {
		routers := h.GetRouter()
		for _, r := range routers {
			method := r.(*Router).Method
			path := r.(*Router).Path
			req := h.CloneReq()
			resp := h.CloneResp()
			engine.Handle(method, path, converHandle(req, resp, beforeHandle, behindHandle, h.Handle))
		}
	}
	return
}

func converHandle(req, resp interface{}, beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handle func(ctx context.Context, req, resp interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = ContextWithRouter(ctx, &Router{
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
		})

		defer func() {
			if r := recover(); r != nil {
				c.Set(HandlerError, HandlerErrorCrash)
				c.Set(HandlerMsg, fmt.Sprintf("%v", r))
				c.Status(http.StatusInternalServerError)
				return
			}
		}()

		// 前置处理
		if beforeHandle != nil {
			if err := beforeHandle(ctx, c, req); err != nil {
				c.Set(HandlerError, HandlerErrorInBefore)
				c.Set(HandlerMsg, err.Error())
				c.Status(http.StatusBadRequest)
				return
			}
		} else {
			req = c.Request
		}

		// 处理
		if err := handle(ctx, req, resp); err != nil {
			c.Set(HandlerError, HandlerErrorInProcess)
			c.Set(HandlerMsg, err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		// 后置处理
		if behindHandle != nil {
			if err := behindHandle(ctx, c, resp); err != nil {
				c.Set(HandlerError, HandlerErrorInBehind)
				c.Set(HandlerMsg, err.Error())
				c.Status(http.StatusInternalServerError)
				return
			}
		} else {
			c.JSON(http.StatusOK, resp)
			return
		}
	}
}
