package web

import (
	"net/http"

	"github.com/serge30/gotraining-gipher/storage"
)

// GifRequest is request payload for storage.Gif data model.
type GifRequest struct {
	storage.Gif
}

// Bind is standard method to implement render.Binder interface.
func (gr *GifRequest) Bind(r *http.Request) error {
	return nil
}
