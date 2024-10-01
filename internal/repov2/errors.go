package repov2

import "errors"

var (
	ErrInternal            = errors.New("internal error")
	ErrUnexpectedBehavior  = errors.New("unexpected behavior")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("permission denied")
	ErrUserIdNotFound      = errors.New("user id not found")
	ErrEventTypeNotFound   = errors.New("event type not found")
	ErrLoginAlreadyInUse   = errors.New("login already in use")
	ErrHostNotFound        = errors.New("host not found")
	ErrObjectAlreadyExists = errors.New("object already exists")
	ErrObjectNotFound      = errors.New("object not found")
	ErrObjectTypeNotFound  = errors.New("object type not found")
)
