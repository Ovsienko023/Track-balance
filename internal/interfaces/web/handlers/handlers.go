package handlers

import (
	"embed"
	"encoding/json"
	"net/http"

	"api/internal/core"
)

type Transport struct {
	Core core.Core
	Fs   *embed.FS
}

func New(c core.Core, fs *embed.FS) *Transport {
	return &Transport{
		Core: c,
		Fs:   fs,
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
