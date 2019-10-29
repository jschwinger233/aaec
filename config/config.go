package config

import (
	"time"

	"github.com/jinzhu/configor"
)

type Config struct {
	Apps []struct {
		Name     string
		OnEvents []struct {
			Kind      string
			Callbacks []struct {
				Command string
				Delay   time.Duration `default:0`
			}
		}
	}
}

var config Config

func New(filename string) Config {
	configor.Load(&config, filename)
	return config
}
