package gatepack

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	uuidplus "github.com/cheetah-fun-gs/goplus/uuid"
	"github.com/cheetah-fun-gs/lego/pkg/so"
)

// JSONPackRouter JSONPackRouter
type JSONPackRouter struct {
	GameID int32 `json:"game_id,omitempty"`
	CMD    int32 `json:"cmd,omitempty"`
}

func (r *JSONPackRouter) String() string {
	return fmt.Sprintf("%d-%d", r.GameID, r.CMD)
}

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
func (pack *JSONPack) GetRouter() so.Router {
	return &JSONPackRouter{
		GameID: pack.GameID,
		CMD:    pack.CMD,
	}
}

// GetLogicPack 获取业务对象
func (pack *JSONPack) GetLogicPack() interface{} {
	return pack.LogicPack
}

// GetContext 获取上下文
func (pack *JSONPack) GetContext() context.Context {
	ctx := context.Background()
	ctx = so.ContextWithRouter(ctx, &struct {
		GameID int32
		CMD    int32
	}{
		GameID: pack.GameID,
		CMD:    pack.CMD,
	})
	ctx = so.ContextWithVersion(ctx, strconv.Itoa(int(pack.Version)))
	ctx = so.ContextWithTraceID(ctx, uuidplus.NewV4().Base62())
	ctx = so.ContextWithSeqID(ctx, int(pack.Seq))
	return ctx
}
