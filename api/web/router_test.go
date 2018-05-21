package web

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/serge30/gotraining-gipher/storage"
	"github.com/stretchr/testify/assert"
)

func getResponse(method string, url string, body io.Reader) (*httptest.ResponseRecorder, storage.Storage, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	store, err := storage.NewFakeStorage()
	if err != nil {
		return nil, nil, err
	}

	rr := httptest.NewRecorder()
	handler := GetRouters(store)
	handler.ServeHTTP(rr, req)

	return rr, store, nil
}

func TestGetRoutersList(t *testing.T) {
	resp, _, err := getResponse("GET", "/gifs", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.Code, "Invalid HTTP status code")
		assert.Contains(t, resp.Body.String(), "Gif1")
	}
}
func TestGetRoutersGet(t *testing.T) {
	resp, _, err := getResponse("GET", "/gifs/1", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.Code, "Invalid HTTP status code")
		assert.Contains(t, resp.Body.String(), "Gif1")
	}
}

func TestGetRoutersGetNotFound(t *testing.T) {
	resp, _, err := getResponse("GET", "/gifs/5", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNotFound, resp.Code, "Invalid HTTP status code")
	}
}

func TestGetRoutersGetBadId(t *testing.T) {
	resp, _, err := getResponse("GET", "/gifs/aaa5", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, resp.Code, "Invalid HTTP status code")
	}
}

func TestGetRoutersCreate(t *testing.T) {
	requestBody := strings.NewReader(`{"name": "test", "slug": "Slug", "width": 300, "height": 200}`)
	resp, store, err := getResponse("POST", "/gifs", requestBody)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, resp.Code, "Invalid HTTP status code")
		assert.Contains(t, resp.Body.String(), "test")

		gif, err := store.GetItem(5)
		assert.NoError(t, err)

		assert.Equal(t, 5, gif.ID)
		assert.Equal(t, "test", gif.Name)
		assert.Equal(t, "Slug", gif.Slug)
		assert.Equal(t, 300, gif.Width)
		assert.Equal(t, 200, gif.Height)
	}
}

func TestGetRoutersCreateInvalid(t *testing.T) {
	requestBody := strings.NewReader(`{"name": "test", "width": 300, "height": 200}`)
	resp, store, err := getResponse("POST", "/gifs", requestBody)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, resp.Code, "Invalid HTTP status code")
		assert.Contains(t, resp.Body.String(), "Invalid request")

		_, err := store.GetItem(5)
		assert.EqualError(t, err, "Id 5 is not found")
	}
}

func TestGetRoutersUpdate(t *testing.T) {
	requestBody := strings.NewReader(`{"name": "test", "width": 300}`)
	resp, store, err := getResponse("PUT", "/gifs/2", requestBody)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.Code, "Invalid HTTP status code")
		assert.Contains(t, resp.Body.String(), "test")

		gif, err := store.GetItem(2)
		assert.NoError(t, err)

		assert.Equal(t, 2, gif.ID)
		assert.Equal(t, "test", gif.Name)
		assert.Equal(t, "gif-2", gif.Slug)
		assert.Equal(t, 300, gif.Width)
		assert.Equal(t, 200, gif.Height)
	}
}

func TestGetRoutersUpdateNotFound(t *testing.T) {
	requestBody := strings.NewReader(`{"name": "test", "width": 300}`)
	resp, store, err := getResponse("PUT", "/gifs/5", requestBody)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNotFound, resp.Code, "Invalid HTTP status code")

		_, err := store.GetItem(5)
		assert.EqualError(t, err, "Id 5 is not found")
	}
}

func TestGetRoutersDelete(t *testing.T) {
	resp, store, err := getResponse("DELETE", "/gifs/2", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, 204, resp.Code, "Invalid HTTP status code")

		_, err := store.GetItem(2)
		assert.EqualError(t, err, "Id 2 is not found")
	}
}

func TestGetRoutersDeleteNotFound(t *testing.T) {
	resp, _, err := getResponse("DELETE", "/gifs/5", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNotFound, resp.Code, "Invalid HTTP status code")
	}
}
