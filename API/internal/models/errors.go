package models

import "errors"

var (
	ErrNotAuthenticated = errors.New("user is not authenticated")
	ErrInvalidJWTToken  = errors.New("invalid jwt-token")
	ErrTraceLayout      = "%s | error: %v"
)
