package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serge30/gotraining-gipher/storage"
	"github.com/stretchr/testify/assert"
)

func TestGetRouters(t *testing.T) {
	req, err := http.NewRequest("GET", "/gifs", nil)
	if err != nil {
		t.Fatal(err)
	}

	store, err := storage.NewFakeStorage()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := GetRouters(store)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Bad")
}
