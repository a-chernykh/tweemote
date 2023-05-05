package server

import (
	"fmt"
	"net/http"
)

type ParamsGetter interface {
	GetRequired(string) string
	Error() error
}

type HTTPParamsGetter struct {
	r   *http.Request
	err error
}

func (pg *HTTPParamsGetter) GetRequired(p string) string {
	query := pg.r.URL.Query()
	val := query.Get(p)
	if val == "" {
		pg.err = fmt.Errorf("%s parameter is required", p)
	}

	return val
}

func (pg *HTTPParamsGetter) Error() error {
	return pg.err
}
