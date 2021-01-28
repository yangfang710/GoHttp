package model

import (
	"context"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

var (
	engine *xorm.Engine

	InitBeans = []interface{}{}
)

func Init(driver string, dsn string, debug bool) error {

	var err error
	engine, err = NewEngine(driver, dsn, debug)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = engine.Ping(); err != nil {
		return errors.WithStack(err)
	}

	for _, bean := range InitBeans {
		if _, err = engine.Where("0").Get(bean); err != nil {
			return errors.Wrapf(err, "check table %s", bean.(xorm.TableName).TableName())
		}
	}

	return nil
}

func GetSession(ctx context.Context) *xorm.Session {
	return engine.NewSession().Context(ctx)
}

func GetEngine() *xorm.Engine {
	return engine
}

func StubEngine(mockedEngine *xorm.Engine) {
	engine = mockedEngine
}
