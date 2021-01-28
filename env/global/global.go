package global

import (
	"time"

	"GoHttp/conf"
)

var (
	Config conf.Config
)

const (
	AppName               = "go-http"
	LoggerTimestampFormat = "2006-01-02 15:04:05.999"
)

var (
	TimeLocation = time.UTC
)
