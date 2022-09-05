package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version    string `yaml:"version"`
	Blockchain struct {
		Token string `yaml:"token"`
	} `yaml:"blockchain"`
	Port string `yaml:"port"`
}

// ParseConfig from config.yml
func ParseConfig(path ...string) Config {
	if len(path) != 1 {
		path = make([]string, 1)
		path[0] = "config.yml"
	}
	c := Config{Version: "1.0", Port: "50005"}

	data, _err := ioutil.ReadFile(path[0])
	if _err != nil {
		log.Printf("config.Get err   #%v ", _err)
	}

	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return c
}
