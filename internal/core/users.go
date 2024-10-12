package core

import "context"

type (
	GetProfileRequest struct{}

	GetProfileResponse struct {
		ID   string
		Name string
	}
)

func (c *Core) GetProfile(ctx context.Context, r *GetProfileRequest) (*GetProfileResponse, error) {

	return &GetProfileResponse{
		ID:   "555",
		Name: "Bob",
	}, nil
}
