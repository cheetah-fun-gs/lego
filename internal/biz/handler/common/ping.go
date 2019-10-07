package common

import (
	"context"

	"github.com/cheetah-fun-gs/goso/internal/biz/handler"
	"github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
	"github.com/cheetah-fun-gs/goso/pkg/so"
)

// 常量定义
var (
	PingNetTypes    = []so.NetType{so.NetTypeNNet, so.NetTypeGNet, so.NetTypeLNet} // 在哪些网络注册
	PingURIS        = []string{"common/ping"}                                      // http 用
	PingHTTPMethods = []string{"POST", "GET"}                                      // http 用
	PingRouters     = sohttp.NewRouters(PingURIS, PingHTTPMethods)
)

// PingReq 请求
type PingReq struct {
	Common *handler.CommonReq `json:"common,omitempty"`
}

// PingResp 响应
type PingResp struct {
	Common *handler.CommonResp `json:"common,omitempty"`
}

// PingHandle 获取服务器时间
func PingHandle(ctx context.Context, req *PingReq, resp *PingResp) error {
	resp.Common = handler.GetcommonRespSuccess()
	return nil
}
