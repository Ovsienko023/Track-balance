package interfaces

import (
	"api/gen/api/apiconnect"
	"api/infrastructure/config"
	"api/internal/core"
	"go.uber.org/zap"
)

type Server struct {
	apiconnect.UnimplementedServiceHandler

	Core *core.Core
}

func New(logger *zap.Logger, cnf *config.Config) (*Server, error) {
	c, err := core.New(logger, cnf, core.Repositories{})
	if err != nil {
		return nil, err
	}

	return &Server{
		Core: c,
	}, nil
}
