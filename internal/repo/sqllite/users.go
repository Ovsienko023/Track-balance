package sqllite

import (
	"context"
	"database/sql"
	"errors"

	"api/internal/repo"
	"api/internal/repo/dbmsg"
)

type UsersRepo struct {
	database *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		database: db,
	}
}

const sqlGetProfile = `
	select id,
	       login,
	       display_name
    from users
    where login = 'admin'`

func (s *UsersRepo) GetProfile(ctx context.Context, msg dbmsg.GetProfile) (*dbmsg.Profile, error) {
	//if err := c.conn(); err != nil {
	//	return nil, err
	//}

	var user dbmsg.Profile

	err := s.database.QueryRowContext(ctx, sqlGetProfile, msg.ID).Scan(
		&user.ID,
		&user.Login,
		&user.DisplayName,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, repo.ErrUserIdNotFound
	case err != nil:
		// TODO analize
		return nil, err
	default:
		// TODO analize
		return &user, nil
	}
}
