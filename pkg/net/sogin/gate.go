package sogin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GateBeforeHandleFunc gate 预处理
func GateBeforeHandleFunc(c *gin.Context, req interface{}) error {
	rawPack, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return nil
	}

	err = json.Unmarshal(rawPack, req)
	if err != nil {
		c.Status(http.StatusBadRequest) // 不建议使用 http code, 可以通过 SetBeforeHandleFunc 替换
		return nil
	}

	return nil
}

// GateAfterHandleFunc gate 结束处理
func GateAfterHandleFunc(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

// GateGetContextFunc gate 不从外部获取 context, 自己生产
func GateGetContextFunc(c *gin.Context) (context.Context, error) {
	ctx := context.Background() // :TODO
	return ctx, nil
}

// NewGate 一个新的gate sogin对象
func NewGate(ports []int) (*SoGin, error) {
	soGin, err := New(ports)
	if err != nil {
		return nil, err
	}
	return soGin, nil
}
