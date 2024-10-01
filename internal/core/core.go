package core

import (
	"go.uber.org/zap"

	"api/infrastructure/config"
	"api/internal/repo/sqllite"
	"api/internal/repov2"
)

type (
	Core struct {
		Logger *zap.Logger
		cnf    *config.Config
		DB     *sqllite.Client

		Repo Repositories
	}

	Repositories struct {
		Users   *repov2.UsersRepo
		Circles *repov2.CirclesRepo
		Areas   *repov2.AreasRepo
	}
)

func New(logger *zap.Logger, cnf *config.Config, repositories Repositories) (*Core, error) {
	return &Core{
		cnf:    cnf,
		Logger: logger,
		//DB:     db,
		Repo: repositories,
	}, nil
}
