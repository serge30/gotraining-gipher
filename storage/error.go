package storage

import (
	"fmt"
)

type NotFoundError int

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Id %d is not found", e)
}

type ValidationError string

func (e ValidationError) Error() string {
	return fmt.Sprintf("Field '%s' is malformed or missing", e)
}
