package dbmsg

type (
	GetCircles struct{}

	Circle struct {
		ID          int64   `json:"id"`
		Areas       []Area  `json:"areas"`
		Description *string `json:"description"`
		CreatedAt   int64   `json:"created_at"`
	}
)

type (
	GetCircle struct {
		CircleID int64 `json:"circle_id"`
	}
)

type (
	CreateCircle struct {
		UserID      int64              `json:"user_id"`
		Areas       []CreateCircleArea `json:"areas"`
		Description *string            `json:"description"`
	}

	CreateCircleArea struct {
		DisplayName string           `json:"display_name"`
		Description *AreaDescription `json:"description"`
		Grade       uint32           `json:"grade"`
	}
)

type DeleteCircle struct {
	ID int64 `json:"id"`
}
