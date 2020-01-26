package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	config CoreConfig = CoreConfig{
		PidFile:     "$HOME/aaec.pid",
		LogLevel:    "DEBUG",
		LogFilename: "$HOME/aaec.log",
	}
)

type CoreConfig struct {
	PidFile     string `yaml:"pidfile"`
	LogLevel    string `yaml:"log_level"`
	LogFilename string `yaml:"log_filename"`
}

func MustInit(filename string) CoreConfig {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("invalid core config filename")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic("invalid core config")
	}

	config.PidFile = os.ExpandEnv(config.PidFile)
	config.LogFilename = os.ExpandEnv(config.LogFilename)
	return config
}

func GetConfig() CoreConfig {
	return config
}
