package gin

import (
	"context"
	"fmt"
	"net/http"

	uuidplus "github.com/cheetah-fun-gs/goplus/uuid"
	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
	"github.com/gin-gonic/gin"
)

// 常量
const (
	LegoRequestID  = "lego-request-id"
	LegoHandlerErr = "lego-handler-err"
	LegoHandlerMsg = "lego-handler-msg"

	HandleCrash       = "handle crash"
	HandleError       = "handle error"
	BeforeHandleError = "before handle error"
	BehindHandleError = "behind handle error"
)

// Register 注册处理器
func Register(engine *gin.Engine, beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handlers ...legocore.Handler) {
	for _, h := range handlers {
		routers := h.GetRouter()
		for _, r := range routers {
			method := r.GetHTTPMethod()
			path := r.GetURI()
			engine.Handle(method, path, converHandle(beforeHandle, behindHandle, h))
		}
	}
	return
}

func converHandle(beforeHandle, behindHandle func(ctx context.Context, c *gin.Context, v interface{}) error, handler legocore.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuidplus.NewV4().Base62()
		ctx := ContextWithValue(&ContextValue{
			RequestID: requestID,
			Path:      c.Request.URL.Path,
			Method:    c.Request.Method,
		})
		c.Set(LegoRequestID, requestID)

		defer func() {
			if r := recover(); r != nil {
				c.Set(LegoHandlerErr, HandleCrash)
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
				c.Set(LegoHandlerErr, BeforeHandleError)
				c.Set(LegoHandlerMsg, err.Error())
				c.Status(http.StatusBadRequest)
				return
			}
		} else {
			req = c.Request
		}

		// 处理
		if err := handler.Handle(ctx, req, resp); err != nil {
			c.Set(LegoHandlerErr, HandleError)
			c.Set(LegoHandlerMsg, err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		// 后置处理
		if behindHandle != nil {
			if err := behindHandle(ctx, c, resp); err != nil {
				c.Set(LegoHandlerErr, BehindHandleError)
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
