package handlers

import (
	"net/http"

	"api/internal/repo/dbmsg"
)

type (
	GetProfileRequest struct{}

	GetProfileResponse struct {
		ID          int64  `json:"id"`
		Login       string `json:"login"`
		DisplayName string `json:"display_name"`
	}
)

// GetProfile ...  Get Profile
// @Summary Get Profile
// @Description Getting user data
// @Tags Profile
// @Param request body GetProfileRequest true "query params"
// @Success 200 {object} GetProfileResponse
// @Router /api/v1/profile [get]
func (t *Transport) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := ErrorResponse{}

	_ = r.Header.Get("Authorization")

	message := dbmsg.GetProfile{}

	result, err := t.Core.Repo.Users.GetProfile(r.Context(), message)
	if err != nil {
		errorContainer.AnalyzeCoreError(w, err)
		return
	}

	response := &GetProfileResponse{
		ID:          result.ID,
		Login:       result.Login,
		DisplayName: result.DisplayName,
	}

	JsonResponse(w, http.StatusOK, response)
}
