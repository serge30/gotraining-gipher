package web

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/serge30/gotraining-gipher/storage"
)

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render mrthod to render ErrResponse type and set correct HTTP response code.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest creates new error response for error "400 Invalid request".
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrRender creates new error response for error "422 Error rendering response".
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound is standard error response for "404 Resource not found".
var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

// ConverStorageError converts storage package errors to renderer.
func ConverStorageError(err error) render.Renderer {
	switch err.(type) {
	case storage.NotFoundError:
		return ErrNotFound
	case storage.ValidationError:
		return ErrInvalidRequest(err)
	}

	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal error.",
		ErrorText:      err.Error(),
	}
}
