package common

import (
	"context"

	"github.com/cheetah-fun-gs/goso/internal/biz/handlers"
	"github.com/cheetah-fun-gs/goso/pkg/so"
)

// 常量定义
var (
	PostNetTypes    = []so.NetType{so.NetTypeNNet, so.NetTypeGNet, so.NetTypeLNet} // 在哪些网络注册
	PostURIS        = []string{"common/post"}                                      // http 用
	PostHTTPMethods = []string{"POST"}                                             // http 用
)

// TestData 测试数据
type TestData struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

// PostReq 请求
type PostReq struct {
	Common *handlers.CommonReq `json:"common,omitempty"`
	Data   *TestData           `json:"data,omitempty"`
}

// PostResp 响应
type PostResp struct {
	Common *handlers.CommonResp `json:"common,omitempty"`
	Data   *TestData            `json:"data,omitempty"`
}

// PostHandle 测试post数据
func PostHandle(ctx context.Context, req *PostReq, resp *PostResp) error {
	resp.Common = handlers.GetcommonRespSuccess()
	resp.Data = req.Data
	return nil
}
