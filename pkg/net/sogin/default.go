package sogin

import (
	"context"
	"encoding/json"
	"goso/pkg/so"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultHook 默认的注入方法
func DefaultHook(req, resp interface{}, handle func(ctx context.Context, req, resp interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawPack, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(rawPack, req)
		if err != nil {
			c.Status(http.StatusGatewayTimeout)
			return
		}

	}
}

// New 一个新的默认sogin对象
func New() (*SoGin, error) {
	router := gin.Default()
	return &SoGin{
		Engine:  router,
		NetAttr: &so.NetAttr{},
	}, nil
}
