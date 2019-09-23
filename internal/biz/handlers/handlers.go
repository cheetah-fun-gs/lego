// Package handlers ...
// 1. 一个处理器一个文件，不要分目录/包
// 2. 处理器逻辑尽量简单，复杂逻辑放到模块包中
package handlers

import (
	"time"
)

// CommonReq 公共请求
type CommonReq struct {
	Version  string `json:"version,omitempty"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Ts       int64  `json:"ts,omitempty"`
}

// CommonRespCode 公共返回码
type CommonRespCode int

// 公共返回码 定义
const (
	CommonRespCodeSuccess CommonRespCode = 0
	CommonRespCodeUnknown                = 1
)

// CommonResp 公共响应
type CommonResp struct {
	Code CommonRespCode `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Ts   int64          `json:"ts,omitempty"`
}

// CommonRespUnknown 位置错误返回
func CommonRespUnknown() *CommonResp {
	return &CommonResp{
		Code: CommonRespCodeUnknown,
		Msg:  "unknown",
		Ts:   time.Now().Unix(),
	}
}

// CommonRespSuccess 成功返回
func CommonRespSuccess() *CommonResp {
	return &CommonResp{
		Code: CommonRespCodeSuccess,
		Msg:  "success",
		Ts:   time.Now().Unix(),
	}
}
