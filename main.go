package main

import (
	flag "github.com/spf13/pflag"

	"github.com/jschwinger23/aaec/app"
	"github.com/jschwinger23/aaec/config"
)

var configFilename string

func init() {
	flag.String(configFilename, "aaec.yaml", "specify the config filename")
}

func main() {
	flag.Parse()
	conf := config.New(configFilename)
	aaec := app.New(conf)
	aaec.Run()
	println("exited")
}
