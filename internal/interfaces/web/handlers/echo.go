package handlers

import (
	"html/template"
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

func (t *Transport) Docs(w http.ResponseWriter, r *http.Request) {
	errorContainer := ErrorResponse{}

	tpl, err := template.ParseFS(t.Fs, "web/apidoc/v1/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := map[string]interface{}{
		"userAgent": r.UserAgent(),
	}

	if err := tpl.Execute(w, data); err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	JsonResponse(w, http.StatusInternalServerError, nil)
}
