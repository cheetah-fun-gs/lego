package common

import (
	"github.com/cheetah-fun-gs/lego/cmd"

	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
)

// 常用全局变量
var (
	GlobalIsDebugMode bool
	GlobalIsAutoInit  bool
	GlobalEnvName     string
)

func init() {
	initConfiger(cmd.EnvName) // 最先

	// 常用全局变量初始化 其次
	GlobalEnvName = cmd.EnvName
	GlobalIsAutoInit = cmd.IsAutoInit
	GlobalIsDebugMode = mconfiger.GetBoolD("is_debug_mode", false)

	if !GlobalIsAutoInit { // 防止scf直接加载所有资源
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
