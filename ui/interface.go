package ui

import (
	"github.com/jschwinger23/aaec/utils"
	"github.com/pkg/errors"
)

var app Application

type Application interface {
	CreateEvent(struct {
		ID        string
		CreatedAt int64
		Package   string
		Type      string
	}) (err error)
}

func InitApplication(a Application) error {
	return utils.DoOnce(func() {
		app = a
	})
}

func GetApplication() (Application, error) {
	if app == nil {
		return nil, errors.New("app not init")
	}
	return app, nil
}
