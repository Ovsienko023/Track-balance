package repov2

import (
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"

	"api/internal/repo"
)

func AnalyzeError(err error) error {
	var dbError sqlite3.Error

	if errors.As(err, &dbError) {
		switch dbError.Code {
		case 19:
			return repo.ErrObjectAlreadyExists
		}
	}

	return fmt.Errorf("%w: %v", repo.ErrInternal, err)
}
