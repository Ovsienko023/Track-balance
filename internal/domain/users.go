package domain

type (
	GetProfile struct {
		ID int64
	}

	Profile struct {
		ID          int64
		Login       string
		DisplayName string
	}
)
