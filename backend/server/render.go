package server

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	ErrorDescription string `json:"error_description"`
}

func renderJson(o interface{}, w http.ResponseWriter) {
	j, err := json.Marshal(o)
	if err != nil {
		renderError(err, w)
		return
	}

	w.Write(j)
}

func renderSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func renderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func renderError(err error, w http.ResponseWriter) {
	http.Error(w, errorJson(err), 500)
}

func errorJson(err error) string {
	var httpError HttpError
	httpError.ErrorDescription = err.Error()
	json, marshalErr := json.Marshal(httpError)
	if marshalErr != nil {
		panic(marshalErr)
	}
	return string(json)
}
