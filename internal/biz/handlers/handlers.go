// Package handlers ...
// 1. 一个处理器一个文件，不要分目录/包
// 2. 处理器逻辑尽量简单，复杂逻辑放到模块包中
package handlers

import (
	"encoding/json"
)

// ReqBody 请求结构体
type ReqBody struct {
	Common ReqCommon       `json:"common,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

// ReqCommon 公共请求
type ReqCommon struct {
	Version  string `json:"version,omitempty"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Ts       int64  `json:"ts,omitempty"`
}

// RespBody 响应结构体
type RespBody struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Ts   int64       `json:"ts,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
