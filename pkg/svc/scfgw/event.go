package scfgw

import (
	"encoding/json"
)

// https://cloud.tencent.com/document/product/583/12513
// {
// 	"requestContext": {
// 		"serviceId": "service-f94sy04v",
// 		"path": "/test/{path}",
// 		"httpMethod": "POST",
// 		"requestId": "c6af9ac6-7b61-11e6-9a41-93e8deadbeef",
// 		"identity": {
// 			"secretId": "abdcdxxxxxxxsdfs"
// 		},
// 		"sourceIp": "10.0.2.14",
// 		"stage": "release"
// 	},
// 	"headers": {
// 		"Accept-Language": "en-US,en,cn",
// 		"Accept": "text/html,application/xml,application/json",
// 		"Host": "service-3ei3tii4-251000691.ap-guangzhou.apigateway.myqloud.com",
// 		"User-Agent": "User Agent String"
// 	},
// 	"body": "{\"test\":\"body\"}",
// 	"pathParameters": {
// 		"path": "value"
// 	},
// 	"queryStringParameters": {
// 		"foo": "bar"
// 	},
// 	"headerParameters": {
// 		"Refer": "10.0.2.14"
// 	},
// 	"stageVariables": {
// 		"stage": "release"
// 	},
// 	"path": "/test/value",
// 	"queryString": {
// 		"foo": "bar",
// 		"bob": "alice"
// 	},
// 	"httpMethod": "POST"
// }

// EventRequestContext ...
type EventRequestContext struct {
	ServiceID  string      `json:"serviceId,omitempty"`
	Path       string      `json:"path,omitempty"`
	HTTPMethod string      `json:"httpMethod,omitempty"`
	RequestID  string      `json:"requestId,omitempty"`
	Identity   interface{} `json:"identity,omitempty"`
	SourceIP   string      `json:"sourceIp,omitempty"`
	Stage      string      `json:"stage,omitempty"`
}

// Event ...
type Event struct {
	RequestContext        *EventRequestContext `json:"requestContext,omitempty"`
	Headers               map[string]string    `json:"headers,omitempty"`
	Body                  json.RawMessage      `json:"body,omitempty"`
	PathParameters        map[string]string    `json:"pathParameters,omitempty"`
	QueryString           map[string]string    `json:"queryString,omitempty"`
	QueryStringParameters map[string]string    `json:"queryStringParameters,omitempty"`
	HeaderParameters      map[string]string    `json:"headerParameters,omitempty"`
	StageVariables        map[string]string    `json:"stageVariables,omitempty"`
	Path                  string               `json:"path,omitempty"`
	HTTPMethod            string               `json:"httpMethod,omitempty"`
	LegoParams            map[string]string    `json:"legoParams,omitempty"` // lego内部使用
}
