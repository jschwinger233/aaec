package main

import (
	"github.com/imroc/req"
	"github.com/pkg/errors"
)

type URI struct {
	endpoint string
	path     string
}

func NewURI(endpoint string) *URI {
	return &URI{endpoint: endpoint}
}

func (u *URI) Compose(subpath string) *URI {
	u.path = u.path + subpath
	return u
}

func (u *URI) Post(object interface{}) (err error) {
	uri := u.endpoint + u.path
	r, err := req.Post(uri, req.BodyJSON(&object))
	if err != nil {
		return errors.Wrap(err, "ui.http.client failed to post events")
	}
	resp := r.Response()
	if resp.StatusCode > 300 {
		return errors.Errorf("ui.http.client http status exception: %v", resp.StatusCode)
	}
	return nil
}
