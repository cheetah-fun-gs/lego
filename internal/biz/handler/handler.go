// Package handler 处理器包
// 1. 一个功能一个子包，只支持一层子包
// 2. 一个处理器一个文件
// 3. 处理器逻辑尽量简单，复杂逻辑放到模块包中
// 4. 不能引用其他处理器包
// 5. 不能引用插件包
package handler

import (
	"fmt"
	"time"

	jsonplus "github.com/cheetah-fun-gs/goplus/encoding/json"
)

// CommonReq 公共请求
type CommonReq struct {
	UID       string `json:"uid,omitempty"`
	Token     string `json:"token,omitempty"`
	Ts        int64  `json:"ts,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

func (commReq *CommonReq) String() string {
	s, _ := jsonplus.ToJSON(commReq)
	return s
}

// CommonRespCode 公共返回码
type CommonRespCode int

// 公共返回码 定义 0 ~ 100 保留, 作为 框架层的返回码
const (
	CommonRespCodeSuccess       CommonRespCode = 0
	CommonRespCodeServerUnknown                = 1 // 服务端未知错误
	CommonRespCodeClientUnknown                = 2 // 客户端未知错误
)

// CommonRespMsg 公共返回文案
var CommonRespMsg = map[CommonRespCode]string{
	CommonRespCodeSuccess:       "success",
	CommonRespCodeServerUnknown: "server unknown",
	CommonRespCodeClientUnknown: "client unknown",
}

// CommonResp 公共响应
type CommonResp struct {
	Code      CommonRespCode `json:"code,omitempty"`
	Msg       string         `json:"msg,omitempty"`
	Ts        int64          `json:"ts,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}

func (commResp *CommonResp) String() string {
	s, _ := jsonplus.ToJSON(commResp)
	return s
}

// GetCommonResp 通用返回
func GetCommonResp(code CommonRespCode, err error) *CommonResp {
	return &CommonResp{
		Code: code,
		Msg:  fmt.Sprintf("%s err: %v", CommonRespMsg[CommonRespCodeSuccess], err),
		Ts:   time.Now().Unix(),
	}
}

// GetcommonRespSuccess 成功返回
func GetcommonRespSuccess() *CommonResp {
	return &CommonResp{
		Code: CommonRespCodeSuccess,
		Msg:  CommonRespMsg[CommonRespCodeSuccess],
		Ts:   time.Now().Unix(),
	}
}

// GetCommonRespServerUnknown 服务端未知错误
func getCommonRespServerUnknown(err error) *CommonResp {
	return GetCommonResp(CommonRespCodeServerUnknown, err)
}
