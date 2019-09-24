package handlers

import (
	"context"
	"github.com/cheetah-fun-gs/goso/pkg/so"
)

// 常量定义
var (
	CommonTimeNetType     = []so.NetType{so.NetTypeGnet, so.NetTypeLnet} // 在哪些网络注册
	CommonTimeURIS        = []string{"common/time"}                      // http 用
	CommonTimeHTTPMethods = []string{"POST", "GET"}                      // http 用
)

// CommonTimeReq 请求
type CommonTimeReq struct {
	Common *CommonReq `json:"common,omitempty"`
}

// CommonTimeResp 响应
type CommonTimeResp struct {
	Common *CommonResp `json:"common,omitempty"`
}

// CommonTimeHandle 获取服务器时间
func CommonTimeHandle(ctx context.Context, req *CommonTimeReq, resp *CommonTimeResp) error {
	resp.Common = getcommonRespSuccess()
	return nil
}
