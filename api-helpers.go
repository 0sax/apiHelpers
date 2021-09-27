package apiHelpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	ErrorResponse struct {
		Message string `json:"message"`
	}

	OKResponse struct {
		Message string `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func WriteErrorJSONResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{message})
}

func WriteOKJSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(OKResponse{
		Message: message,
		Data:    data,
	})
}

func Unmarshal(body io.ReadCloser, target interface{}) error {

	bd, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bd, target); err != nil {
		return err
	}

	return nil
}
