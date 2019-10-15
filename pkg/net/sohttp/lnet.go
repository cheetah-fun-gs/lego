package sohttp

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/cheetah-fun-gs/goso/pkg/utils"
	"github.com/gin-gonic/gin"
)

func lNetPreHandleFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, int, error) {
	data := map[string]interface{}{}
	for key, val := range c.Request.PostForm {
		if strings.HasPrefix(key, ContextPrefix) {
			if len(val) == 1 {
				data[strings.Replace(key, ContextPrefix, "", 1)] = val[0]
			}
		}
	}

	ctx := utils.LoadContext(data, so.ContextRevert)

	rawPack, err := c.GetRawData()
	if err != nil {
		soHTTP.Logger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return ctx, http.StatusBadRequest, err
	}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, req)
		if err != nil {
			soHTTP.Logger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return ctx, http.StatusBadRequest, err
		}
	}
	return ctx, http.StatusOK, nil
}

// NewLNet 一个新的lnet gin 对象
func NewLNet() (*SoHTTP, error) {
	lnet, err := New()
	if err != nil {
		return nil, err
	}

	if err := lnet.SetBeforeHandleFunc(lNetPreHandleFunc); err != nil {
		return nil, err
	}

	return lnet, nil
}
