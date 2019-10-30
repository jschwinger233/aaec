package ui

import (
	"errors"
	"sync"
)

var once sync.Once
var app Application

type Application interface {
	ReportEvent(Event) Result
	TraceEvent(EventId) Result
}

func InitApp(a Application) error {
	init := false
	once.Do(func() {
		init = true
		app = a
	})
	if !init {
		return errors.New("init app more than once")
	}
	return nil
}

func GetApp() (Application, error) {
	if app == nil {
		return nil, errors.New("app not init")
	}
	return app, nil
}
