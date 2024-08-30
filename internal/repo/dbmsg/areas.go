package dbmsg

type (
	GetBaseAreas struct {
		UserID int64
	}

	Area struct {
		ID          int64            `json:"id"`
		DisplayName string           `json:"display_name"`
		Description *AreaDescription `json:"description,omitempty"`
		Grade       uint32           `json:"grade"`
	}

	AreaDescription struct {
		Progress string `json:"progress"`
		Target   string `json:"target"`
	}
)
