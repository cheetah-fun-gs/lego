package handlers

import "context"

// CommonTimeReq 请求
type CommonTimeReq struct {
}

// CommonTimeResp 响应
type CommonTimeResp struct {
}

// 常量定义
const (
	CommonTimeURI        = "common/time"
	CommonTimeHTTPMethod = "POST"
)

// CommonTime 获取服务器时间
func CommonTime(ctx context.Context, req *CommonTimeReq, resp *CommonTimeResp) error {
	return nil
}
