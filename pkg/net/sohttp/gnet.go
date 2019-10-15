package sohttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/gatepack"
	"github.com/gin-gonic/gin"
)

func gNetPreHandleFunc(soHTTP *SoHTTP, c *gin.Context, req interface{}) (context.Context, int, error) {
	ctx := context.Background()

	gatePack := soHTTP.GatePack

	rawPack, err := c.GetRawData()
	if err != nil {
		soHTTP.Logger.Error(ctx, "BadRequest GetRawData error: %v", err)
		return nil, http.StatusBadRequest, err
	}

	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, gatePack)
		if err != nil {
			soHTTP.Logger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return nil, http.StatusBadRequest, err
		}
	}

	err = gatePack.Verify()
	if err != nil {
		soHTTP.Logger.Error(ctx, "BadRequest Verify error: %v", err)
		return nil, http.StatusBadRequest, err
	}

	ctx = gatePack.GetContext()

	logicPack := gatePack.GetLogicPack()
	if logicPack != nil {
		err = json.Unmarshal(logicPack.(json.RawMessage), req)
		if err != nil {
			soHTTP.Logger.Error(ctx, "BadRequest Unmarshal error: %v", err)
			return nil, http.StatusBadRequest, err
		}
	}
	return ctx, http.StatusOK, nil
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

	if err := gnet.SetBeforeHandleFunc(gNetPreHandleFunc); err != nil {
		return nil, err
	}

	return gnet, nil
}
