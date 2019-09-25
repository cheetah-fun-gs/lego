package sohttp

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cheetah-fun-gs/goso/pkg/utils"
	"github.com/gin-gonic/gin"
)

func lNetUnmarshalFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, error) {
	data := map[string]interface{}{}
	for key, val := range c.Request.PostForm {
		if strings.HasPrefix(key, ContextPrefix) {
			if len(val) == 1 {
				data[strings.Replace(key, ContextPrefix, "", 1)] = val[0]
			}
		}
	}
	ctx := utils.LoadContext(data)

	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return ctx, err
	}
	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, req)
		if err != nil {
			soLogger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return ctx, err
		}
	}
	return ctx, nil
}

// NewLNet 一个新的lnet gin 对象
func NewLNet() (*SoHTTP, error) {
	lnet, err := New()
	if err != nil {
		return nil, err
	}

	if err := lnet.SetUnmarshalFunc(lNetUnmarshalFunc); err != nil {
		return nil, err
	}

	return lnet, nil
}
