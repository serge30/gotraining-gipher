package storage

import (
	"fmt"
)

// NotFoundError is returned when requested Id is not in storage.
type NotFoundError int

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Id %d is not found", int(e))
}

// ValidationError is returned when Gif is not valid.
type ValidationError string

func (e ValidationError) Error() string {
	return fmt.Sprintf("Field '%s' is malformed or missing", string(e))
}
