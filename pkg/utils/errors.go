package utils

import "errors"

var (
	ErrEmptyID        = errors.New("ID cannot be empty")
	ErrEmptyTitle     = errors.New("title cannot be empty")
	ErrEmptyStatus    = errors.New("status cannot be empty")
	ErrInvalidStatus  = errors.New("the status is invalid")
	ErrInvalidId      = errors.New("the id is invalid")
	ErrTaskNotFound   = errors.New("task not found")
	ErrInvalidPayload = errors.New("invalid request payload")
	ErrFailedEncode   = errors.New("failed to encode task")
)
