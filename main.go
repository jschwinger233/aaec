package main

import (
	flag "github.com/spf13/pflag"

	"github.com/jschwinger23/aaec/app"
	"github.com/jschwinger23/aaec/config"
	"github.com/jschwinger23/aaec/logging"
)

func main() {
	coreConfigFilename, appConfigFilename := parseFlags()
	config.MustInit(coreConfigFilename)
	logging.MustInit()
	aaec := app.New(appConfigFilename)
	aaec.Run()
	logger := logging.GetLogger("main")
	logger.Errorf("exited")
}

func parseFlags() (coreConfigFilename, appConfigFilename string) {
	flag.StringVar(&coreConfigFilename, "core-config", "config.example.yaml", "specify the core config filename")
	flag.StringVar(&appConfigFilename, "app-config", "app.example.yaml", "specify the app config filename")
	flag.Parse()
	return
}
