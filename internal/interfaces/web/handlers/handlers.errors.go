package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"api/internal/core"
)

const (
	InvalidRequest = "invalid request"
)

type ErrorResponseDetails struct {
	Reason      string `json:"reason,omitempty"`
	Description string `json:"description,omitempty"`
	Position    string `json:"position,omitempty"`
}

type ErrorResponseError struct {
	Code        int                    `json:"code,omitempty"`
	Description string                 `json:"description,omitempty"`
	Details     []ErrorResponseDetails `json:"details"`
}

type ErrorResponse struct {
	Error ErrorResponseError `json:"error,omitempty"`
}

func (r *ErrorResponse) Add(reason, description, position string) {
	details := ErrorResponseDetails{
		Reason:      reason,
		Description: description,
		Position:    position,
	}
	r.Error.Details = append(r.Error.Details, details)
}

func (r *ErrorResponse) Done(w http.ResponseWriter, code int, description string) {
	r.Error.Code = code
	r.Error.Description = description
	if len(r.Error.Details) == 0 {
		r.Error.Details = []ErrorResponseDetails{}
	}

	response, _ := json.Marshal(r)

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write([]byte(response))

}

func (r *ErrorResponse) Marshaling() ([]byte, error) {
	if len(r.Error.Details) == 0 {
		r.Error.Details = []ErrorResponseDetails{}
	}

	response, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *ErrorResponse) AnalyzeCoreError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, core.ErrUnauthorized):
		r.Done(w, http.StatusUnauthorized, err.Error())
		return

	case errors.Is(err, core.ErrPermissionDenied):
		r.Done(w, http.StatusForbidden, err.Error())
		return

	case errors.Is(err, core.ErrObjectNotFound), errors.Is(err, core.ErrObjectTypeNotFound):
		r.Done(w, http.StatusNotFound, err.Error())
		return

	case errors.Is(err, core.ErrObjectAlreadyExists):
		r.Done(w, http.StatusConflict, err.Error())
		return

	case err != nil:
		r.Done(w, http.StatusInternalServerError, "internal error")
		return

	}
}
