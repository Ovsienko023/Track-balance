package interfaces

import (
	"api/gen/api/domain"
	"api/internal/core"
)

func GetProfileRequestToCore(r *domain.GetProfileRequest) *core.GetProfileRequest {
	return &core.GetProfileRequest{}
}

func GetProfileResponseFromCore(r *core.GetProfileResponse) *domain.GetProfileResponse {
	return &domain.GetProfileResponse{
		Id:   r.ID,
		Name: r.Name,
	}
}
