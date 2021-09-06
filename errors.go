package apiHelpers

import "net/http"

type ClientErr interface {
	ClientReadable() bool
	ClientMessage() string
	ClientStatusCode() int
}

func IsClientReadable(err error) bool {
	ce, ok := err.(ClientErr)
	return ok && ce.ClientReadable()
}

func WriteError(w http.ResponseWriter, err error, elseStatus int, elseMessage string) {
	if IsClientReadable(err) {
		e := err.(ClientErr)
		WriteErrorJSONResponse(w, e.ClientStatusCode(), e.ClientMessage())
		return
	}
	WriteErrorJSONResponse(w, elseStatus, elseMessage)
	return
}