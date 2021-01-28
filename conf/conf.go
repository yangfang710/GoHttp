package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	Config struct {
		Mode string `json:"mode" env:"APP_ENV"`

		Addr string `json:"addr"`

		Metrics struct {
			Addr string `json:"addr"`
		} `json:"metrics"`

		DebugOption struct {
			GlobalEnable bool `json:"global_enable"`
		} `json:"debug"`

		Cors struct {
			AllowAllOrigins bool     `json:"allowAllOrigins"`
			AllowOrigins    []string `json:"allowOrigins"`
			AllowMethods    []string `json:"allowMethods"`
			AllowHeaders    []string `json:"allowHeaders"`
		} `json:"cors"`

		DB struct {
			Driver string `json:"driver"`
			DSN    string `json:"dsn" env:"DB_MEOW"`
		} `json:"db"`

		TimeZone string `json:"time_zone" env:"TIME_ZONE"`
	}
)

func (c *Config) Debug() bool {
	return c.DebugOption.GlobalEnable
}

// Decode decode json data to *Config
func Decode(data []byte) (*Config, error) {
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

// ParseFile parse Config from config file and system env
func ParseFile(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// Parse parse Config from json data and system env
func Parse(data []byte) (*Config, error) {
	c, err := Decode(data)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func ParseTestFile() (*Config, error) {
	var (
		projectRoot string
		path        string
	)

	if projectRoot = os.Getenv("PROJECT_ROOT"); projectRoot == "" {
		projectRoot = fmt.Sprintf("%s/src/github.com/yangfang710/GoHttp", os.Getenv("GOPATH"))
	}
	path = fmt.Sprintf("%s/conf/app.test.json", projectRoot)

	c, err := ParseFile(path)
	if err != nil {
		return nil, err
	}

	return c, nil
}
