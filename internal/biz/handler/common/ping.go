package common

import (
	"context"

	"github.com/cheetah-fun-gs/lego/internal/biz/handler"
)

// 常量定义
var (
	SvcNames        = []string{}              // 在哪些服务注册 默认全部
	PingURIS        = []string{"common/ping"} // http 必填
	PingHTTPMethods = []string{"POST", "GET"} // http 必填
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
