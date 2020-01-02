package main

import (
	// 顺序不能错
	_ "github.com/cheetah-fun-gs/lego/cmd"
	_ "github.com/cheetah-fun-gs/lego/internal/common"
	_ "github.com/cheetah-fun-gs/lego/internal/svc"
)

//go:generate legoctl project gen $GOFILE
func main() {
}
