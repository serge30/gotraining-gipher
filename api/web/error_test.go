package web

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serge30/gotraining-gipher/storage"

	"github.com/go-chi/render"
	"github.com/stretchr/testify/assert"
)

func renderResponse(v render.Renderer) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	err = render.Render(rr, req, v)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

func TestErrNotFound(t *testing.T) {
	resp := ErrNotFound

	rr, err := renderResponse(resp)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNotFound, rr.Code, "Invalid HTTP status code")
		assert.Contains(t, rr.Body.String(), "Resource not found")
	}
}

func TestErrInvalidRequest(t *testing.T) {
	err := errors.New("test")
	resp := ErrInvalidRequest(err)

	rr, err := renderResponse(resp)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Invalid HTTP status code")
		assert.Contains(t, rr.Body.String(), "Invalid request")
		assert.Contains(t, rr.Body.String(), "test")
	}
}

func TestErrRender(t *testing.T) {
	err := errors.New("test")
	resp := ErrRender(err)

	rr, err := renderResponse(resp)
	if assert.NoError(t, err) {
		assert.Equal(t, 422, rr.Code, "Invalid HTTP status code")
		assert.Contains(t, rr.Body.String(), "Error rendering response")
		assert.Contains(t, rr.Body.String(), "test")
	}
}

func TestConverStorageErrorNotFound(t *testing.T) {
	err := storage.NotFoundError(5)
	resp := ConverStorageError(err)

	assert.IsType(t, (*ErrResponse)(nil), resp)
	errResponse, ok := resp.(*ErrResponse)
	assert.True(t, ok)

	assert.Equal(t, ErrNotFound, errResponse)
}

func TestConverStorageErrorValidation(t *testing.T) {
	err := storage.ValidationError("test")
	resp := ConverStorageError(err)

	assert.IsType(t, (*ErrResponse)(nil), resp)
	errResponse, ok := resp.(*ErrResponse)
	assert.True(t, ok)

	assert.Equal(t, 400, errResponse.HTTPStatusCode)
	assert.Equal(t, "Invalid request.", errResponse.StatusText)
	assert.Equal(t, "Field 'test' is malformed or missing", errResponse.ErrorText)
}

func TestConverStorageErrorOther(t *testing.T) {
	err := errors.New("test")
	resp := ConverStorageError(err)

	assert.IsType(t, (*ErrResponse)(nil), resp)
	errResponse, ok := resp.(*ErrResponse)
	assert.True(t, ok)

	assert.Equal(t, 500, errResponse.HTTPStatusCode)
	assert.Equal(t, "Internal error.", errResponse.StatusText)
	assert.Equal(t, "test", errResponse.ErrorText)
}
