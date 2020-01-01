package common

import (
	"fmt"
	"time"

	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	mmgodb "github.com/cheetah-fun-gs/goplus/multier/multimgodb"
	"github.com/cheetah-fun-gs/goplus/structure"
	"github.com/globalsign/mgo"
)

func initMgo(dbs map[string]interface{}) {
	if v, ok := dbs["default"]; !ok {
		panic("mgo dbs.default not configuration")
	} else {
		// 初始化默认 mgo 连接池
		dbConfig := &MgoConfig{}
		if err := structure.MapToStruct(v.(map[string]interface{}), dbConfig); err != nil {
			panic(err)
		}

		defaultDB, err := dbConfig.Conn()
		if err != nil {
			panic(err)
		}
		mmgodb.Init(defaultDB)
	}

	// 初始化其他 mgo 连接池
	for dbName, dbData := range dbs {
		if dbName != "default" {
			dbConfig := &MgoConfig{}
			if err := structure.MapToStruct(dbData.(map[string]interface{}), dbConfig); err != nil {
				panic(err)
			}

			db, err := dbConfig.Conn()
			if err != nil {
				panic(err)
			}

			if err := mmgodb.Register(dbName, db); err != nil {
				panic(err)
			}
		}
	}
}

// MgoConfig mongoDB configuration
type MgoConfig struct {
	Addr                string `yml:"addr,omitempty" json:"addr,omitempty"`
	AuthSource          string `yml:"auth_source,omitempty" json:"auth_source,omitempty"`
	Username            string `yml:"username,omitempty" json:"username,omitempty"`
	Password            string `yml:"password,omitempty" json:"password,omitempty"`
	DB                  string `yml:"db,omitempty" json:"db,omitempty"`
	PoolLimit           int    `yml:"pool_limit,omitempty" json:"pool_limit,omitempty"`
	PoolTimeout         int    `yml:"pool_timeout,omitempty" json:"pool_timeout,omitempty"`
	MaxIdleTime         int    `yml:"max_idle_time,omitempty" json:"max_idle_time,omitempty"`
	Timeout             int    `yml:"timeout,omitempty" json:"timeout,omitempty"`
	IsUseTestCollection bool   `yml:"is_use_test_collection,omitempty" json:"is_use_test_collection,omitempty"` // 是否在 collection name 后加 .test  一个业务需求
}

// Conn Conn
func (m *MgoConfig) Conn() (*mgo.Database, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:         []string{m.Addr},
		Username:      m.Username,
		Password:      m.Password,
		Source:        m.AuthSource,
		Database:      m.DB,
		PoolLimit:     m.PoolLimit,
		PoolTimeout:   time.Duration(m.PoolTimeout) * time.Second,
		MaxIdleTimeMS: m.MaxIdleTime * 1000,
		Timeout:       time.Duration(m.Timeout) * time.Second,
	}
	if dialInfo.PoolLimit == 0 {
		dialInfo.PoolLimit = 10 // 连接池限制一定要设置
	}
	if dialInfo.PoolTimeout == 0 {
		dialInfo.PoolTimeout = 2 * time.Second // 从连接池获取连接的超时时间一定要设置
	}
	if dialInfo.MaxIdleTimeMS == 0 {
		dialInfo.MaxIdleTimeMS = 60 * 1000 // 连接池释放空闲连接的时间一定要设置
	}
	if dialInfo.Timeout == 0 {
		dialInfo.Timeout = 5 * time.Second // 连接超时时间一定要设置
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return session.DB(m.DB), nil
}

// CollectionName 获取 默认 mongo collection name
func CollectionName(collectName string) string {
	if mconfiger.GetBoolDN("mongo", "dbs.default.is_use_test_collection", false) {
		return collectName + ".test"
	}
	return collectName
}

// CollectionNameN 获取 指定 mongo collection name
func CollectionNameN(name, collectName string) string {
	key := fmt.Sprintf("dbs.%s.is_use_test_collection", name)
	if mconfiger.GetBoolDN("mongo", key, false) {
		return collectName + ".test"
	}
	return collectName
}
