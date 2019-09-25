// Package handlers ...
// 1. 一个处理器一个文件，不要分目录/包
// 2. 处理器逻辑尽量简单，复杂逻辑放到模块包中
package handlers

import (
	"fmt"
	"time"

	"github.com/cheetah-fun-gs/goso/pkg/so"
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

// 公共返回码 定义 0 ~ 100 保留, 作为 框架层的返回码
const (
	CommonRespCodeSuccess       CommonRespCode = 0
	CommonRespCodeServerUnknown                = 1 // 服务端未知错误 BadGateway
	CommonRespCodeClientUnknown                = 2 // 客户端未知错误 BadRequest
)

// CommonResp 公共响应
type CommonResp struct {
	Code CommonRespCode `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Ts   int64          `json:"ts,omitempty"`
}

// getCommonRespUnknown 未知错误返回
func getCommonRespUnknown(code CommonRespCode, err error) *CommonResp {
	return &CommonResp{
		Code: code,
		Msg:  fmt.Sprintf("%v", err),
		Ts:   time.Now().Unix(),
	}
}

// getcommonRespSuccess 成功返回
func getcommonRespSuccess() *CommonResp {
	return &CommonResp{
		Code: CommonRespCodeSuccess,
		Msg:  "success",
		Ts:   time.Now().Unix(),
	}
}

// getCommonRespServerUnknown 服务端未知错误
func getCommonRespServerUnknown(err error) *CommonResp {
	return getCommonRespUnknown(CommonRespCodeServerUnknown, err)
}

// HandleCommonRespSoNet 框架层错误处理
func HandleCommonRespSoNet(code so.ErrorNetCode, err error) interface{} {
	switch code {
	case so.ErrorNetCodeBadRequest:
		return getCommonRespUnknown(CommonRespCodeClientUnknown, err)
	default:
		return getCommonRespUnknown(CommonRespCodeServerUnknown, err)
	}
}
