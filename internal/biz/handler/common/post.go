package common

import (
	"context"

	"github.com/cheetah-fun-gs/lego/internal/biz/handler"
	"github.com/cheetah-fun-gs/lego/pkg/net/sohttp"
	"github.com/cheetah-fun-gs/lego/pkg/so"
)

// 常量定义
var (
	PostNetTypes    = []so.NetType{so.NetTypeNNet, so.NetTypeGNet, so.NetTypeLNet} // 在哪些网络注册
	PostURIS        = []string{"common/post"}                                      // http 用
	PostHTTPMethods = []string{"POST"}                                             // http 用
	PostRouters     = sohttp.NewRouters(PostURIS, PostHTTPMethods)
)

// TestData 测试数据
type TestData struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

// PostReq 请求
type PostReq struct {
	Common *handler.CommonReq `json:"common,omitempty"`
	Data   *TestData          `json:"data,omitempty"`
}

// PostResp 响应
type PostResp struct {
	Common *handler.CommonResp `json:"common,omitempty"`
	Data   *TestData           `json:"data,omitempty"`
}

// PostHandle 测试post数据
func PostHandle(ctx context.Context, req *PostReq, resp *PostResp) error {
	resp.Common = handler.GetcommonRespSuccess()
	resp.Data = req.Data
	return nil
}
