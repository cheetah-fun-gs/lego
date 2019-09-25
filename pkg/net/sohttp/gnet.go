package sohttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheetah-fun-gs/goso/pkg/gatepack"

	"github.com/cheetah-fun-gs/goso/pkg/so"
	"github.com/gin-gonic/gin"
)

// gnetBeforeHandleFunc
func gnetParseRequest(c *gin.Context, gatePack so.GatePack, req interface{}) (ctx context.Context, err error) {
	rawPack, err := c.GetRawData()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest GetRawData error: %v", err)
		return nil, err
	}

	if len(rawPack) != 0 {
		err = json.Unmarshal(rawPack, gatePack)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			return nil, err
		}
	}

	err = gatePack.Verify()
	if err != nil {
		soLogger.Error(context.Background(), "BadRequest Verify error: %v", err)
		return nil, err
	}

	logicPack := gatePack.GetLogicPack()
	if logicPack != nil {
		err = json.Unmarshal(logicPack.([]byte), req)
		if err != nil {
			soLogger.Error(context.Background(), "BadRequest Unmarshal error: %v", err)
			return nil, err
		}
	}

	ctx = gatePack.GetContext()
	return ctx, nil
}

func gnetConverFunc(soHTTP *SoHTTP, handler so.Handler) gin.HandlerFunc {
	req := handler.GetReq()
	resp := handler.GetResp()

	return func(c *gin.Context) {
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				soLogger.Error(ctx, "BadGateway gnetConverFunc error: %v", r)
				errorHandle(c, soHTTP, http.StatusBadGateway, fmt.Errorf("%v", r))
				return
			}
		}()

		ctx, err := gnetParseRequest(c, soHTTP.Config.GatePack, req)
		if err != nil {
			soLogger.Error(ctx, "BadGateway gnetParseRequest error: %v", err)
			errorHandle(c, soHTTP, http.StatusBadRequest, err)
			return
		}

		err = handler.Handle(ctx, req, resp)
		if err != nil {
			soLogger.Error(ctx, "BadGateway %v Handle error: %v", handler.GetName(), err)
			errorHandle(c, soHTTP, http.StatusBadGateway, err)
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// GNet gnet http 对象
type GNet struct {
	*SoHTTP
}

// SetGatePack 设置包对象
func (gnet *GNet) SetGatePack(gatePack so.GatePack) error {
	gnet.SoHTTP.Config.GatePack = gatePack
	return nil
}

// NewGNet 一个新的 gnet http 对象
// http 协议的 gnet 没啥实际价值， 请用专业的 api 网关
func NewGNet() (so.GNet, error) {
	soHTTP, err := New()
	if err != nil {
		return nil, err
	}
	if err := soHTTP.SetConverFunc(gnetConverFunc); err != nil {
		return nil, err
	}

	gnet := &GNet{
		SoHTTP: soHTTP,
	}

	gatePack := &gatepack.JSONPack{}
	if err := gnet.SetGatePack(gatePack); err != nil {
		return nil, err
	}
	return gnet, nil
}
