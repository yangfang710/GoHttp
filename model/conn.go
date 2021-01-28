package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var (
	maxOpenConns    = 32
	maxIdleConns    = 32
	connMaxLifeTime = 8 * time.Hour
	logLevel        = core.LOG_ERR
)

func NewEngine(driver, dsn string, debug bool) (*xorm.Engine, error) {

	engine, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		return nil, err
	}

	engine.SetMaxOpenConns(maxOpenConns)
	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetConnMaxLifetime(connMaxLifeTime)
	engine.SetLogLevel(logLevel)
	if debug {
		engine.ShowSQL(true)
	}

	return engine, nil
}
