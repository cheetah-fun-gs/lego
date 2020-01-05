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

	if !GlobalIsAutoInit {
		return
	}

	initLogger() // 其次

	// 存在对应配置则初始化对应模块

	if ok, dbs, err := mconfiger.GetMapN("redigo", "dbs"); err != nil {
		panic(err)
	} else if ok {
		initRedigo(dbs)
	}

	if ok, dbs, err := mconfiger.GetMapN("sql", "dbs"); err != nil {
		panic(err)
	} else if ok {
		initSQLDB(dbs)
	}

	if ok, dbs, err := mconfiger.GetMapN("mgo", "dbs"); err != nil {
		panic(err)
	} else if ok {
		initMgo(dbs)
	}
}
