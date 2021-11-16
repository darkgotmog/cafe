package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Config struct {
	Port int
}

func LoadConfig(path string) (Config, error) {

	conf := Config{}

	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return conf, err
	}

	conf.Port = cfg.Section("server").Key("port").MustInt(1323)
	return conf, nil
}
