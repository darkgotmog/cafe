package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Config struct {
	Port int
}

func LoadConfig(path string) (Config, error) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("patch:", exPath)

	conf := Config{}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Fail to current dir %v", err)
		return conf, err
	}
	fmt.Println("dir:", dir)

	cfg, err := ini.Load(dir + "/" + path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return conf, err
	}

	conf.Port = cfg.Section("server").Key("port").MustInt(1323)
	return conf, nil
}
