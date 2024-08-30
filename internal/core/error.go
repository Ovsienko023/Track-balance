package core

import "errors"

var (
	ErrInternal            = errors.New("Internal server error. ")
	ErrUnauthorized        = errors.New("Unauthorized. ")
	ErrPermissionDenied    = errors.New("Permission denied. ")
	ErrInvalidArguments    = errors.New("Invalid argument. ")
	ErrServerConnection    = errors.New("Could not connect to the server. ")
	ErrUserNotFound        = errors.New("User not found. ")
	ErrObjectAlreadyExists = errors.New("Object already exist. ")
	ErrObjectNotFound      = errors.New("Object not found. ")
	ErrObjectTypeNotFound  = errors.New("Object type not found. ")
)
