package common

import (
	"time"

	mredigopool "github.com/cheetah-fun-gs/goplus/multier/multiredigopool"
	"github.com/cheetah-fun-gs/goplus/structure"
	redigo "github.com/gomodule/redigo/redis"
)

func initRedigo(dbs map[string]interface{}) {
	if v, ok := dbs["default"]; !ok {
		panic("redigo dbs.default not configuration")
	} else {
		// 初始化默认redis连接池
		dbConfig := &RedigoConfig{}
		if err := structure.MapToStruct(v.(map[string]interface{}), dbConfig); err != nil {
			panic(err)
		}
		mredigopool.Init(dbConfig.Pool())
	}

	// 初始化其他redis连接池
	for dbName, dbData := range dbs {
		if dbName != "default" {
			dbConfig := &RedigoConfig{}
			if err := structure.MapToStruct(dbData.(map[string]interface{}), dbConfig); err != nil {
				panic(err)
			}
			if err := mredigopool.Register(dbName, dbConfig.Pool()); err != nil {
				panic(err)
			}
		}
	}
}

// RedigoConfig RedigoConfig
type RedigoConfig struct {
	Addr        string `yml:"addr,omitempty" json:"addr,omitempty"`
	Auth        string `yml:"auth,omitempty" json:"auth,omitempty"`
	DB          int    `yml:"db,omitempty" json:"db,omitempty"`
	MaxIdle     int    `yml:"max_idle,omitempty" json:"max_idle,omitempty"` // pool
	MaxActive   int    `yml:"max_active,omitempty" json:"max_active,omitempty"`
	IdleTimeout int    `yml:"idle_timeout,omitempty" json:"idle_timeout,omitempty"`
	PingRate    int    `yml:"ping_rate,omitempty" json:"ping_rate,omitempty"` // ping 的 频率 单位 /秒
}

// Conn Conn
func (redigoConfig *RedigoConfig) Conn() (redigo.Conn, error) {
	c, err := redigo.Dial("tcp", redigoConfig.Addr)
	if err != nil {
		return nil, err
	}
	if redigoConfig.Auth != "" {
		if _, err := c.Do("AUTH", redigoConfig.Auth); err != nil {
			c.Close()
			return nil, err
		}
	}
	if _, err := c.Do("SELECT", redigoConfig.DB); err != nil {
		c.Close()
		return nil, err
	}
	return c, err
}

// Pool Pool
func (redigoConfig *RedigoConfig) Pool() *redigo.Pool {
	if redigoConfig.MaxIdle == 0 {
		redigoConfig.MaxIdle = 5
	}
	if redigoConfig.MaxActive == 0 {
		redigoConfig.MaxActive = 10
	}
	if redigoConfig.IdleTimeout == 0 {
		redigoConfig.IdleTimeout = 60
	}
	return &redigo.Pool{
		MaxIdle:     redigoConfig.MaxIdle,
		MaxActive:   redigoConfig.MaxActive,
		IdleTimeout: time.Duration(redigoConfig.IdleTimeout) * time.Second,
		Wait:        true,
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if redigoConfig.PingRate == 0 {
				return nil
			}
			if time.Since(t) < time.Duration(redigoConfig.PingRate)*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Dial: redigoConfig.Conn,
	}
}
