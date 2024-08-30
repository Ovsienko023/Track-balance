package sqllite

import (
	"context"
	"database/sql"

	"api/internal/repo/dbmsg"
)

type AreasRepo struct {
	database *sql.DB
}

func NewAreasRepo(db *sql.DB) *AreasRepo {
	return &AreasRepo{
		database: db,
	}
}

func (l *AreasRepo) GetBaseAreas(ctx context.Context, msg dbmsg.GetBaseAreas) ([]dbmsg.Area, error) {
	//TODO implement me
	panic("implement me")
}
