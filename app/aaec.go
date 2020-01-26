package app

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/juju/loggo"
)

var (
	logger log.Logger
)

type AAEC struct {
}

func init() {
	logger = log.GetLogger("app")
}

func New(filename string) *AAEC {
	return &AAEC{}
}

func (a *AAEC) Run() {
	go a.installSigalHandlers()
}

func (a *AAEC) stop() {
}

func (a *AAEC) wait() {
}

func (a *AAEC) reload() {
}

func (a *AAEC) installSigalHandlers() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)
	signal.Notify(ch, syscall.SIGHUP)

	for {
		switch sig := <-ch; sig {
		case syscall.SIGINT, syscall.SIGTERM:
			logger.Infof("got signal 2 or 15 to gracefully exit")
			a.stop()
			a.wait()
			return

		case syscall.SIGHUP:
			logger.Infof("got signal 1 to reload config")
			a.reload()
			continue

		default:
			logger.Infof("got signal and ignore:", sig)
			continue
		}
	}

}
