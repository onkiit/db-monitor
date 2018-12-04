package helper

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	Error error `json:"error"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	res := &ErrResponse{
		Error: err,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
