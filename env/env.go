package env

import (
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"GoHttp/env/global"
	"GoHttp/model"
)

func logger() error {

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: global.LoggerTimestampFormat,
	}
	// enable debug for non production env
	if global.Config.Debug() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(formatter)

	return nil
}

func databases() error {
	return model.Init(global.Config.DB.Driver, global.Config.DB.DSN, global.Config.Debug())
}

func timeLocation() error {
	loc, err := time.LoadLocation(global.Config.TimeZone)
	if err != nil {
		return errors.WithStack(err)
	}

	global.TimeLocation = loc
	return nil
}
