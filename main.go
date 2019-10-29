package main

import (
	log "github.com/juju/loggo"
	flag "github.com/spf13/pflag"

	"github.com/jschwinger23/aaec/app"
	"github.com/jschwinger23/aaec/config"
)

var (
	configFilename string
	logger         log.Logger
)

func init() {
	flag.StringVar(&configFilename, "config", "config.example.yaml", "specify the config filename")
	logger = log.GetLogger("main")
}

func main() {
	flag.Parse()
	conf := config.New(configFilename)
	aaec := app.New(conf)
	aaec.Run()
	logger.Infof("exited")
}
