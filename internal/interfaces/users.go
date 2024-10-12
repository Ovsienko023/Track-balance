package interfaces

import (
	"api/gen/api/domain"
	"connectrpc.com/connect"
	"context"
)

func (s *Server) GetProfile(
	ctx context.Context,
	req *connect.Request[domain.GetProfileRequest],
) (*connect.Response[domain.GetProfileResponse], error) {
	res, err := s.Core.GetProfile(ctx, GetProfileRequestToCore(req.Msg))
	if err != nil {
		return nil, AnalyzeError(err)
	}

	return connect.NewResponse(GetProfileResponseFromCore(res)), nil
}
