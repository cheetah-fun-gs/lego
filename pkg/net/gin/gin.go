package gin

import (
	"context"
	"net/http"

	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
	"github.com/gin-gonic/gin"
)

// 常量
const (
	ErrorBadRequest   = "BadRequest"
	ErrorHandlerCrash = "HandlerCrash"
)

// Register 注册处理器
func Register(engine *gin.Engine, beforeHandle, behindHandle func(c *gin.Context, v interface{}) error, handlers ...legocore.Handler) {
	for _, h := range handlers {
		routers := h.GetRouter()
		for _, r := range routers {
			httpMethod := r.(*Router).HTTPMethod
			uri := r.(*Router).URI
			req := h.CloneReq()
			resp := h.CloneResp()
			engine.Handle(httpMethod, uri, converHandle(req, resp, beforeHandle, behindHandle, h.Handle))
		}
	}
	return
}

func converHandle(req, resp interface{}, beforeHandle, behindHandle func(c *gin.Context, v interface{}) error, handle func(ctx context.Context, req, resp interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = ContextWithRouter(ctx, &Router{
			HTTPMethod: c.Request.Method,
			URI:        c.Request.URL.Path,
		})

		defer func() {
			if r := recover(); r != nil {
				c.Set(ErrorHandlerCrash, r)
				c.Status(http.StatusInternalServerError)
				return
			}
		}()

		// 前置处理
		var err error
		if beforeHandle != nil {
			err = beforeHandle(c, req)
		} else {
			req, err = c.GetRawData()
		}

		if err != nil {
			c.Set(ErrorBadRequest, err)
			c.Status(http.StatusBadRequest)
			return
		}

		// 处理
		if err = handle(ctx, req, resp); err != nil {
			c.Set(ErrorHandlerCrash, err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// 后置处理
		if behindHandle != nil {
			if err = behindHandle(c, resp); err != nil {
				c.Set(ErrorHandlerCrash, err)
				c.Status(http.StatusInternalServerError)
				return
			}
		} else {
			c.JSON(http.StatusOK, resp)
			return
		}
	}
}
