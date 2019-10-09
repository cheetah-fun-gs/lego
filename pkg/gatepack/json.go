package gatepack

import (
	"context"
	"encoding/json"

	"github.com/cheetah-fun-gs/goso/pkg/so"
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
	ctx = context.WithValue(ctx, &so.ContextTraceID{}, uuid.NewV4().String())
	ctx = context.WithValue(ctx, &so.ContextVersion{}, pack.Version)
	ctx = context.WithValue(ctx, &so.ContextSeqID{}, pack.Seq)
	ctx = context.WithValue(ctx, &so.ContextRouter{}, &struct {
		GameID int32
		CMD    int32
	}{
		GameID: pack.GameID,
		CMD:    pack.CMD,
	})
	return ctx
}
