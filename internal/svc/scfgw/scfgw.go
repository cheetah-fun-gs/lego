package scfgw

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/cheetah-fun-gs/lego/internal/biz/handler"
	legoscfgw "github.com/cheetah-fun-gs/lego/pkg/svc/scfgw"
)

// 常量
const (
	LegoRequestID = "lego-request-id"
)

// BeforeHandleFunc ...
func BeforeHandleFunc(ctx context.Context, event *legoscfgw.Event, req interface{}) error {
	if err := json.Unmarshal(event.Body, req); err != nil {
		return err
	}

	// 获取公共请求头
	commonReq := reflect.ValueOf(req).Elem().FieldByName("Common").Interface().(*handler.CommonReq)
	// 合并客户端request id
	requestids := []string{}
	if commonReq.RequestID != "" {
		requestids = append(requestids, commonReq.RequestID)
	}
	requestids = append(requestids, event.RequestContext.RequestID)

	commonReq.RequestID = strings.Join(requestids, "--")
	event.LegoParams[LegoRequestID] = commonReq.RequestID
	return nil
}

// BehindHandleFunc ...
func BehindHandleFunc(ctx context.Context, event *legoscfgw.Event, resp interface{}) error {
	// 获取公共响应头
	commonResp := reflect.ValueOf(resp).Elem().FieldByName("Common").Interface().(*handler.CommonResp)
	commonResp.RequestID = event.LegoParams[LegoRequestID]
	return nil
}
