package errors

import "errors"

var (
	ErrInvalidGroupData = errors.New("invalid group data")
	ErrGroupNotFound    = errors.New("Group not found")
)
