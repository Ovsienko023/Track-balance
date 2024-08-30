package handlers

import (
	"encoding/json"
	"net/http"

	"api/internal/core"
)

type Transport struct {
	Core core.Core
}

func New(c core.Core) *Transport {
	return &Transport{
		Core: c,
	}
}

func JsonResponse(w http.ResponseWriter, code int, resp any) {
	if resp == nil {
		w.WriteHeader(code)
		return
	}

	response, err := json.Marshal(resp)
	if err != nil {
		errorContainer := ErrorResponse{}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func FileResponse(w http.ResponseWriter, file []byte, filename string) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	_, err := w.Write(file)
	if err != nil {
		return err
	}

	return nil
}
