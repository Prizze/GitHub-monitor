package errors

import (
	"errors"
)

var (
	ErrLanguageUnexpected  = errors.New("language is unexpected")
	ErrFailedCreateRequest = errors.New("failed to create request")
	ErrFailedRequest       = errors.New("request failed")
)
