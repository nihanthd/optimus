package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nihanthd/optimus/pi"

	"github.com/nihanthd/optimus/env"
	"github.com/nihanthd/optimus/log"
	"github.com/nihanthd/optimus/metrics/statsd"
	"github.com/nihanthd/optimus/pins/handler"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Env    env.Config     `yaml:"env"`
		Log    log.Config     `yaml:"log"`
		Http   handler.Config `yaml:"http"`
		Statsd statsd.Config  `yaml:"statsd"`
		Pi     pi.Config      `yaml:"pi"`
	}
)

var (
	verbose    bool
	configFlag string
	config     *Config
)

var (
	rootCmd = &cobra.Command{
		Use:     "optimus",
		Short:   "Optimus is a middleware REST server for a Raspberry Pi",
		Long:    `Optimus is a middleware REST server for a Raspberry Pi`,
		Version: "0.0.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			c, err := NewConfig(configFlag)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			config = c
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configFlag, "config", "c", "config.yaml", "Config file path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Configuration
func NewConfig(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Reading of config file failed %s", path)
	}
	config := &Config{}

	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, errors.Wrapf(err, "Unmarshal of config file failed %s", path)
	}
	fmt.Printf("%+v\n", config)
	return config, nil
}
