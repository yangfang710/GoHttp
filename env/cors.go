package env

import (
	"github.com/gin-contrib/cors"

	"GoHttp/env/global"
)

func CorsConfig() cors.Config {
	c := global.Config.Cors
	config := cors.DefaultConfig()

	if c.AllowAllOrigins {
		config.AllowAllOrigins = true
	} else {
		// 过滤域名
	}

	config.AllowHeaders = c.AllowHeaders
	config.AllowMethods = c.AllowMethods

	return config
}
