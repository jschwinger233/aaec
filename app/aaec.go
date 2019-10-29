package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jschwinger23/aaec/config"
)

type AAEC struct {
}

func New(config config.Config) *AAEC {
	return &AAEC{}
}

func (a *AAEC) Run() {
	a.installSigalHandler()
}

func (a *AAEC) stop() {
}

func (a *AAEC) wait() {
}

func (a *AAEC) reload() {
}

func (a *AAEC) installSigalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)
	signal.Notify(ch, syscall.SIGHUP)

	for {
		switch sig := <-ch; sig {
		case syscall.SIGINT, syscall.SIGTERM:
			println("got signal 2 or 15 to graceful exit")
			a.stop()
			a.wait()
			return

		case syscall.SIGHUP:
			println("got signal 1 to reload config")
			a.reload()

		default:
			println("got signal and ignore:", sig)
			continue
		}
	}

}
