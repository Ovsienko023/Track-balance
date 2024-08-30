package handlers

import (
	"net/http"
)

type (
	EchoRequest struct{}

	EchoResponse struct {
		Code string `json:"code"`
	}
)

func (t *Transport) Echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	_ = r.Header.Get("Authorization")

	response := EchoResponse{
		Code: "OK",
	}

	JsonResponse(w, http.StatusOK, response)
}
