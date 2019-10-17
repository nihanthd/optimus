package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Environment string   `yaml:"environment"`
	Server      *Server  `yaml:"server"`
	Logging     *Logging `yaml:"logging"`
}

type Server struct {
	HTTPPort      int64 `yaml:"http_port"`
	HTTPSPort     int64 `yaml:"https_port"`
	WebsocketPort int64 `yaml:"websocket_port"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Access bool   `yaml:"access"`
}

/*
Reads the yaml file and loads into the Config struct
*/
func Parse(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Reading of config file failed %s", path)
	}
	config := &Config{}

	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, errors.Wrapf(err, "Unmarshal of config file failed %s", path)
	}
	return config, nil
}
