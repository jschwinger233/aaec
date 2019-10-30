package ui

import (
	log "github.com/juju/loggo"
)

var logger log.Logger

func init() {
	logger = log.GetLogger("ui")
}
