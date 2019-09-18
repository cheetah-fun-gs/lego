package sogin

import (
	"context"
	"encoding/json"
	"goso/pkg/so"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultBeforeHandleFunc 默认的预处理
func DefaultBeforeHandleFunc(c *gin.Context, req interface{}) (ctx context.Context, err error) {
	rawPack, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return
	}

	err = json.Unmarshal(rawPack, req)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return
	}

	return ctx, nil
}

// DefaultAfterHandleFunc 默认的
func DefaultAfterHandleFunc(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

// New 一个新的默认sogin对象
func New() (*SoGin, error) {
	router := gin.Default()
	return &SoGin{
		Engine: router,
		NetAttr: &so.NetAttr{
			Ports: []int{8000},
		},
		BeforeHandleFunc: DefaultBeforeHandleFunc,
		AfterHandleFunc:  DefaultAfterHandleFunc,
	}, nil
}
