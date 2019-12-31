package common

import (
	"database/sql"
	"fmt"

	sqlplus "github.com/cheetah-fun-gs/goplus/dao/sql"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	msqldb "github.com/cheetah-fun-gs/goplus/multier/multisqldb"
	"github.com/cheetah-fun-gs/goplus/structure"
	"github.com/go-sql-driver/mysql"
)

func initSQLDB() {
	// 初始化默认sql连接池
	defaultSQLConfig := &SQLConfig{}

	ok, err := mconfiger.GetStructN("sql", "dbs.default", defaultSQLConfig)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("sql dbs.default not configuration")
	}

	defaultDB, err := defaultSQLConfig.Open()
	if err != nil {
		panic(err)
	}

	in := sqlplus.NewSafeInterceptor()
	msqldb.InitWithInterceptor(defaultDB, in)

	// 初始化其他sql连接池
	_, dbs, _ := mconfiger.GetMapN("sql", "dbs")
	for dbName, dbData := range dbs {
		if dbName != "default" {
			dbConfig := &SQLConfig{}
			if err := structure.MapToStruct(dbData.(map[string]interface{}), dbConfig); err != nil {
				panic(err)
			}
			db, err := dbConfig.Open()
			if err != nil {
				panic(err)
			}

			in := sqlplus.NewSafeInterceptor()
			if err := msqldb.RegisterWithInterceptor(dbName, db, in); err != nil {
				panic(err)
			}
		}
	}
}

// SQLConfig SQLConfig
type SQLConfig struct {
	DriverName   string `yml:"driver_name,omitempty" json:"driver_name,omitempty"`
	MaxOpenConns int    `yml:"max_open_conns,omitempty" json:"max_open_conns,omitempty"`
	MaxIdleConns int    `yml:"max_idle_conns,omitempty" json:"max_idle_conns,omitempty"`
	User         string `yml:"user,omitempty" json:"user,omitempty"`       // Username
	Passwd       string `yml:"passwd,omitempty" json:"passwd,omitempty"`   // Password (requires User)
	Net          string `yml:"net,omitempty" json:"net,omitempty"`         // Network type
	Addr         string `yml:"addr,omitempty" json:"addr,omitempty"`       // Network address (requires Net)
	DBName       string `yml:"db_name,omitempty" json:"db_name,omitempty"` // Database name
}

// ToDSN ToDSN
func (sqlConfig *SQLConfig) ToDSN() string {
	switch sqlConfig.DriverName {
	case "mysql":
		cfg := &mysql.Config{
			User:   sqlConfig.User,
			Passwd: sqlConfig.Passwd,
			Net:    sqlConfig.Net,
			Addr:   sqlConfig.Addr,
			DBName: sqlConfig.DBName,
			Params: map[string]string{"charset": "utf8mb4", "parseTime": "true", "loc": "Asia/Shanghai"},
		}
		return cfg.FormatDSN()
	default:
		panic(fmt.Errorf("DriverName %s is not support", sqlConfig.DriverName))
	}
}

// Open Open
func (sqlConfig *SQLConfig) Open() (*sql.DB, error) {
	switch sqlConfig.DriverName {
	case "mysql":
		dsn := sqlConfig.ToDSN()
		db, err := sql.Open(sqlConfig.DriverName, dsn)
		if err != nil {
			return nil, err
		}
		if sqlConfig.MaxIdleConns != 0 {
			db.SetMaxIdleConns(sqlConfig.MaxIdleConns)
		}
		if sqlConfig.MaxOpenConns != 0 {
			db.SetMaxOpenConns(sqlConfig.MaxOpenConns)
		}
		return db, nil
	default:
		return nil, fmt.Errorf("DriverName %s is not support", sqlConfig.DriverName)
	}
}
