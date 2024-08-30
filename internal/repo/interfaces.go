package repo

import (
	"context"

	"api/internal/repo/dbmsg"
)

type (
	UserRepo interface {
		GetProfile(ctx context.Context, msg dbmsg.GetProfile) (*dbmsg.Profile, error)
	}

	CirclesRepo interface {
		SearchCircles(ctx context.Context, msg dbmsg.GetCircles) ([]dbmsg.Circle, error)
		CreateCircle(ctx context.Context, msg dbmsg.CreateCircle) (*int64, error)
	}

	LabelsRepo interface {
		GetBaseAreas(ctx context.Context, msg dbmsg.GetBaseAreas) ([]dbmsg.Area, error)
	}
)
