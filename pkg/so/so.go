// Package so goso 各接口的定义
package so

import (
	"fmt"
)

// Addr 地址
type Addr struct {
	IP   string
	Port int
}

func (addr *Addr) String() string {
	return fmt.Sprintf("%s:%d", addr.IP, addr.Port)
}

// NetType 网络类型
type NetType int

// 网络类型
const (
	NetTypeGnet NetType = iota
	NetTypeLnet
)
