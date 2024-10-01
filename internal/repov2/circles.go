package repov2

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"api/internal/repo/dbmsg"
)

type CirclesRepo struct {
	database *sql.DB
}

func NewCircles(db *sql.DB) *CirclesRepo {
	return &CirclesRepo{
		database: db,
	}
}

const getCirclesSQL = `
	select 
	    id,
	    description,
	    created_at
	from circles
`

func (c *CirclesRepo) SearchCircles(ctx context.Context, msg dbmsg.GetCircles) ([]dbmsg.Circle, error) {
	//if err := c.conn(); err != nil {
	//	return nil, err
	//}

	rows, err := c.database.QueryContext(ctx, getCirclesSQL)
	if err != nil {
		return nil, AnalyzeError(err)
	}

	objects := make([]dbmsg.Circle, 0)

	for rows.Next() {
		item := dbmsg.Circle{}
		err := rows.Scan(
			&item.ID,
			&item.Description,
			&item.CreatedAt,
		)

		if err != nil {
			// TODO log
			continue
		}

		area, _ := c.getAreasCircle(ctx, item.ID)
		if len(area) > 0 {
			item.Areas = area
		}

		objects = append(objects, item)
	}

	return objects, nil
}

const getCircleSQL = `
	select 
	    id,
	    description,
	    created_at
	from circles
	where id = $1
`

func (c *CirclesRepo) GetCircle(ctx context.Context, msg dbmsg.GetCircle) (*dbmsg.Circle, error) {
	//if err := c.conn(); err != nil {
	//	return nil, err
	//}

	rows, err := c.database.QueryContext(ctx, getCircleSQL, msg.CircleID)
	if err != nil {
		return nil, AnalyzeError(err)
	}

	rows.Next()

	object := dbmsg.Circle{}
	err = rows.Scan(
		&object.ID,
		&object.Description,
		&object.CreatedAt,
	)

	if err != nil {
		// TODO log
		return nil, AnalyzeError(err)
	}

	area, _ := c.getAreasCircle(ctx, object.ID)
	if len(area) > 0 {
		object.Areas = area
	}

	return &object, nil
}

const getAreasCircleSQL = `
	select 
	    id,
	    display_name,
	    description,
	    grade
	from areas
	where circle_id = $1
`

func (c *CirclesRepo) getAreasCircle(ctx context.Context, circleID int64) ([]dbmsg.Area, error) {
	rows, err := c.database.QueryContext(ctx, getAreasCircleSQL, circleID)
	if err != nil {
		return nil, AnalyzeError(err)
	}

	objects := make([]dbmsg.Area, 0)

	for rows.Next() {
		item := dbmsg.Area{}
		var description []byte

		err := rows.Scan(
			&item.ID,
			&item.DisplayName,
			&description,
			&item.Grade,
		)

		if err != nil {
			// TODO log
			continue
		}

		if len(description) > 0 {
			err = json.Unmarshal(description, &item.Description)
			if err != nil {
				// TODO log
				continue
			}
		}

		objects = append(objects, item)
	}

	return objects, nil
}

const createCircleSQL = `
	insert into circles (user_id, description, created_at) 
	values (?, ?, ?)
`

func (c *CirclesRepo) CreateCircle(ctx context.Context, msg dbmsg.CreateCircle) (*int64, error) {
	rows, err := c.database.ExecContext(ctx, createCircleSQL,
		msg.UserID,
		msg.Description,
		time.Now().Unix(),
	)

	if err != nil {
		return nil, AnalyzeError(err)
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, AnalyzeError(err)
	}

	for _, area := range msg.Areas {
		if _, err = c.createArea(ctx, area, msg.UserID, id); err != nil {
			// TODO: log
		}
	}

	return &id, nil
}

const createAreaSQL = `
	insert into areas (user_id, circle_id, display_name, description, grade) 
	values (?, ?, ?, ?, ?)
`

func (c *CirclesRepo) createArea(ctx context.Context, msg dbmsg.CreateCircleArea, userID int64, circleID int64) (*int64, error) {
	description, _ := json.Marshal(msg.Description)

	rows, err := c.database.ExecContext(ctx, createAreaSQL,
		userID,
		circleID,
		msg.DisplayName,
		description,
		msg.Grade,
	)

	if err != nil {
		return nil, AnalyzeError(err)
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, AnalyzeError(err)
	}

	return &id, nil
}

const deleteCircleSQL = `
	delete
	from circles
	where id = $1
`

func (c *CirclesRepo) DeleteCircle(ctx context.Context, msg dbmsg.DeleteCircle) error {
	rows, err := c.database.ExecContext(ctx, deleteCircleSQL,
		msg.ID,
	)

	if err != nil {
		return AnalyzeError(err)
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return AnalyzeError(err)
	}

	return nil
}
