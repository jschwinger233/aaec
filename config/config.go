package config

import (
	"io/ioutil"
	"strings"
	"time"

	log "github.com/juju/loggo"
	"gopkg.in/yaml.v3"
)

var (
	logger log.Logger
	config Config = Config{
		LogLevel: "INFO",
	}
)

type Config struct {
	LogLevel string `yaml:"log_level"`

	Apps []struct {
		Name     string
		OnEvents []struct {
			Kind      string
			Callbacks []struct {
				Command string
				Delay   time.Duration
			}
		} `yaml:"on_events"`
	}
}

func init() {
	logger = log.GetLogger("config")
}

func New(filename string) Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Criticalf("invalid config filename: %v", err)
		panic("invalid config filename")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		logger.Criticalf("invalid config: %v", err)
		panic("invalid config")
	}

	var logLevel log.Level
	switch l := strings.ToLower(config.LogLevel); l {
	case "debug":
		logLevel = log.DEBUG
	case "info":
		logLevel = log.INFO
	case "error":
		logLevel = log.ERROR
	case "critical":
		logLevel = log.CRITICAL
	}
	rootLogger := log.GetLogger("")
	rootLogger.SetLogLevel(logLevel)
	return config
}
