package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"api/internal/repo/dbmsg"
)

// GetCircle ... Get Circle
// @Summary Get Circle
// @Description
// @Tags Circles
// @Param   circle_id   path      string  true  "circle_id"
// @Success 200 {object} dbmsg.Circle
// @Router /api/v1/circle/{circle_id} [get]
func (t *Transport) GetCircle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := ErrorResponse{}

	_ = r.Header.Get("Authorization")

	circleId := chi.URLParam(r, "circle_id")
	if circleId == "" {
		errorContainer.Done(w, http.StatusNotFound, errors.New("object not found").Error())
		return
	}

	id, err := strconv.ParseInt(circleId, 10, 64)
	if err != nil {
		errorContainer.Done(w, http.StatusBadRequest, errors.New("invalid circle_id").Error())
		return
	}

	message := dbmsg.GetCircle{
		CircleID: id,
	}

	circle, err := t.Core.DB.CirclesRepo.GetCircle(r.Context(), message)
	if err != nil {
		errorContainer.AnalyzeCoreError(w, err)
		return
	}

	//result := struct {
	//}{}

	JsonResponse(w, http.StatusOK, circle)
}

type (
	SearchCirclesRequest struct{}

	SearchCirclesResponse struct {
		Circles []dbmsg.Circle `json:"circles"`
	}
)

// SearchCircles ... Get all circles
// @Summary Get all circles
// @Description get all circles
// @Tags Circles
// @Param request body SearchCirclesRequest true "query params"
// @Success 200 {object} SearchCirclesResponse
// @Router /api/v1/circles [get]
func (t *Transport) SearchCircles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := ErrorResponse{}

	_ = r.Header.Get("Authorization")

	message := dbmsg.GetCircles{}

	circles, err := t.Core.DB.CirclesRepo.SearchCircles(r.Context(), message)
	if err != nil {
		errorContainer.AnalyzeCoreError(w, err)
		return
	}

	response := &SearchCirclesResponse{
		Circles: circles,
	}

	JsonResponse(w, http.StatusOK, response)
}

type (
	CreateCircleRequest struct {
		UserID      int64                    `json:"user_id"`
		Areas       []dbmsg.CreateCircleArea `json:"areas"`
		Description *string                  `json:"description"`
	}

	CreateCircleResponse struct {
		ID int64 `json:"id"`
	}
)

// CreateCircle ...  Create circle
// @Summary Create circle
// @Description Create circle
// @Tags Circles
// @Param request body CreateCircleRequest true "body params"
// @Success 201 {object} CreateCircleResponse
// @Router /api/v1/circle [post]
func (t *Transport) CreateCircle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := ErrorResponse{}

	decoder := json.NewDecoder(r.Body)
	var message dbmsg.CreateCircle

	err := decoder.Decode(&message)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	_ = r.Header.Get("Authorization")

	// TODO: delete
	profile, err := t.Core.DB.UsersRepo.GetProfile(r.Context(), dbmsg.GetProfile{})
	if err != nil {
		errorContainer.AnalyzeCoreError(w, err)
		return
	}

	message.UserID = profile.ID

	id, err := t.Core.DB.CirclesRepo.CreateCircle(r.Context(), message)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := CreateCircleResponse{
		ID: *id,
	}

	JsonResponse(w, http.StatusCreated, response)
}

// DeleteCircle ...  Delete circle
// @Summary Delete circle
// @Description Delete circle
// @Tags Circles
// @Param   id   path      string  true  "circle_id"
// @Success 204
// @Router /api/v1/circle/{circle_id} [delete]
func (t *Transport) DeleteCircle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := ErrorResponse{}

	circleId := chi.URLParam(r, "circle_id")
	if circleId == "" {
		errorContainer.Done(w, http.StatusNotFound, errors.New("object not found").Error())
		return
	}

	id, err := strconv.ParseInt(circleId, 10, 64)
	if err != nil {
		errorContainer.Done(w, http.StatusBadRequest, errors.New("invalid circle_id").Error())
		return
	}

	message := dbmsg.DeleteCircle{
		ID: id,
	}

	if err := t.Core.DB.CirclesRepo.DeleteCircle(r.Context(), message); err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	JsonResponse(w, http.StatusNoContent, nil)
}
