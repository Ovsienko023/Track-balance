package interfaces

import (
	"api/internal/core"
	"connectrpc.com/connect"
	"errors"
)

func AnalyzeError(err error) error {
	switch {
	case errors.Is(core.ErrUserNotFound, err):
		return connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewError(connect.CodeUnknown, err)
}
