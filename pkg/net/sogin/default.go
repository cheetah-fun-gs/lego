package sogin

import (
	"context"
	"encoding/json"
	"goso/pkg/so"
	"goso/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// DefaultBeforeHandleFunc 默认的预处理
func DefaultBeforeHandleFunc(c *gin.Context, req interface{}) error {
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

// DefaultAfterHandleFunc 默认的结束处理
func DefaultAfterHandleFunc(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

// DefaultGetContextFunc 默认的获取 ctx 的方法
func DefaultGetContextFunc(c *gin.Context) (context.Context, error) {
	data := map[string]interface{}{}
	for _, p := range c.Params {
		if strings.HasPrefix(p.Key, ContextPrefix) {
			data[strings.Replace(p.Key, ContextPrefix, "", 1)] = p.Value
		}
	}
	return utils.LoadContext(data), nil
}

// New 一个新的默认sogin对象
func New() (*SoGin, error) {
	router := gin.Default()
	return &SoGin{
		Engine: router,
		NetAttr: &so.NetAttr{
			Ports: []int{8000},
		},
		GetContextFunc:   DefaultGetContextFunc,
		BeforeHandleFunc: DefaultBeforeHandleFunc,
		AfterHandleFunc:  DefaultAfterHandleFunc,
	}, nil
}
