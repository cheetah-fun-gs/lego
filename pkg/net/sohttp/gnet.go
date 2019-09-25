package sohttp

import (
	"context"
	"encoding/json"

	"github.com/cheetah-fun-gs/goso/pkg/gatepack"
	"github.com/gin-gonic/gin"
)

func gNetUnmarshalFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, error) {
	ctx := context.Background()

	gatePack := soHTTP.GatePack

	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return nil, err
	}

	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, gatePack)
		if err != nil {
			soLogger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return nil, err
		}
	}

	err = gatePack.Verify()
	if err != nil {
		soLogger.Error(ctx, "BadRequest Verify error: %v", err)
		return nil, err
	}

	ctx = gatePack.GetContext()

	logicPack := gatePack.GetLogicPack()
	if logicPack != nil {
		err = json.Unmarshal(logicPack.(json.RawMessage), req)
		if err != nil {
			soLogger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return nil, err
		}
	}
	return ctx, nil
}

// NewGNet 一个新的 gnet http 对象
// http 协议的 gnet 没啥实际价值， 请用专业的 api 网关
func NewGNet() (*SoHTTP, error) {
	gnet, err := New()
	if err != nil {
		return nil, err
	}

	if err := gnet.SetGatePack(&gatepack.JSONPack{}); err != nil {
		return nil, err
	}

	if err := gnet.SetUnmarshalFunc(gNetUnmarshalFunc); err != nil {
		return nil, err
	}

	return gnet, nil
}
