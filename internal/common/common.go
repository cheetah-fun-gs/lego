package common

import (
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
)

// StartType 启动类型
type StartType int

// 常量
const (
	StartTypeWeb StartType = iota
	StartTypeScfgw
)

// Opt ...
type Opt struct {
	EnvName     string
	StartType   StartType
	IsDebugMode bool
}

// 常用全局变量
var (
	GlobalIsDebugMode bool
	GlobalEnvName     string
)

// Init common 不可以有 init 方法, 统一放在 这里
func Init(opt *Opt) {
	initConfiger(opt.EnvName) // 最先

	// 常用全局变量初始化 其次
	GlobalEnvName = opt.EnvName
	GlobalIsDebugMode = mconfiger.GetBoolD("is_debug_mode", false)

	if opt.StartType != StartTypeWeb { // 防止scf直接加载所有资源
		return
	}

	InitLogger(ParseLogger()) // 其次

	// 存在对应配置则初始化对应模块
	redigos := ParseRedigo()
	if len(redigos) > 0 {
		InitRedigo(redigos)
	}

	sqls := ParseSQLDB()
	if len(sqls) > 0 {
		InitSQLDB(sqls)
	}

	mgos := ParseMgo()
	if len(mgos) > 0 {
		InitMgo(mgos)
	}
}
