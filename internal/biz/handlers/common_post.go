package handlers

import (
	"context"
	"goso/pkg/so"
)

// 常量定义
var (
	CommonPostNetType     = []so.NetType{so.NetTypeGnet, so.NetTypeLnet} // 在哪些网络注册
	CommonPostURIS        = []string{"common/post"}                      // http 用
	CommonPostHTTPMethods = []string{"POST"}                             // http 用
)

// TestData 测试数据
type TestData struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

// CommonPostReq 请求
type CommonPostReq struct {
	Common *CommonReq `json:"common,omitempty"`
	Data   *TestData  `json:"data,omitempty"`
}

// CommonPostResp 响应
type CommonPostResp struct {
	Common *CommonResp `json:"common,omitempty"`
	Data   *TestData   `json:"data,omitempty"`
}

// CommonPostHandle 获取服务器时间
func CommonPostHandle(ctx context.Context, req *CommonPostReq, resp *CommonPostResp) error {
	resp.Common = commonRespOK()
	resp.Data = req.Data
	return nil
}