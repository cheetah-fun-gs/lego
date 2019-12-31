package common

import (
	"fmt"
	"os"
	"path/filepath"

	viperplus "github.com/cheetah-fun-gs/goplus/configer/viper"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

// initConfiger 初始化配置
func initConfiger(envName string) {
	if envName == "" {
		envName = "local"
	}

	// 定位配置文件根目录
	absPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(filepath.Dir(absPath), "configs", envName)
	ok, err := filepathplus.Exists(configDir)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(fmt.Sprintf("%s not found in %s", envName, configDir))
	}

	// 默认配置初始化
	defaultViper, err := viperplus.New("default", configDir)
	if err != nil {
		panic(err)
	}
	mconfiger.Init(defaultViper)

	// 所有配置初始化
	filePaths, err := filepathplus.Files(configDir)
	if err != nil {
		panic(err)
	}
	for _, fPath := range filePaths {
		configName := filepathplus.NoExt(filepath.Base(fPath))
		if configName != "default" {
			vip, err := viperplus.New(configName, filepath.Dir(fPath))
			if err != nil {
				panic(err)
			}
			if err := mconfiger.Register(configName, vip); err != nil {
				panic(err)
			}
		}
	}
}
