package core

import (
	"go.uber.org/zap"

	"api/infrastructure/config"
	"api/internal/repo/sqllite"
)

type (
	Core struct {
		Logger *zap.Logger
		cnf    *config.Config
		DB     *sqllite.Client

		repo Repositories
	}

	Repositories struct{}
)

func New(logger *zap.Logger, cnf *config.Config, db *sqllite.Client) (*Core, error) {
	return &Core{
		cnf:    cnf,
		Logger: logger,
		DB:     db,
		repo:   Repositories{},
	}, nil
}
