package gatepack

import (
	"context"
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// JSONPack  json 格式的 gate 包
type JSONPack struct {
	Version   int16           `json:"version,omitempty"`
	GameID    int32           `json:"game_id,omitempty"`
	CMD       int32           `json:"cmd,omitempty"`
	Seq       int32           `json:"seq,omitempty"`
	LogicPack json.RawMessage `json:"logic_pack,omitempty"`
}

// Verify 校验包的有效性
func (pack *JSONPack) Verify() error {
	return nil
}

// GetRouter 获取路由
func (pack *JSONPack) GetRouter() interface{} {
	return nil
}

// GetLogicPack 获取业务对象
func (pack *JSONPack) GetLogicPack() interface{} {
	return pack.LogicPack
}

// GetContext 获取上下文
func (pack *JSONPack) GetContext() context.Context {
	ctx := context.Background()
	context.WithValue(ctx, ContextKey("version"), pack.Version)
	context.WithValue(ctx, ContextKey("game_id"), pack.GameID)
	context.WithValue(ctx, ContextKey("cmd"), pack.CMD)
	context.WithValue(ctx, ContextKey("seq"), pack.Seq)
	context.WithValue(ctx, ContextKey("trace_id"), fmt.Sprintf("%v", uuid.NewV4()))
	return ctx
}
