package utils

import (
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

var onces map[reflect.Value]*sync.Once

func DoOnce(f func()) error {
	funcValue := reflect.ValueOf(f)
	once, ok := onces[funcValue]
	if !ok {
		once = &sync.Once{}
		onces[funcValue] = once
	}

	firstTime := false
	once.Do(func() {
		firstTime = true
		f()
	})

	if !firstTime {
		return errors.New("function executed more than once")
	}
	return nil
}
