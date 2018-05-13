package web

import (
	"net/http"

	"github.com/serge30/gotraining-gipher/storage"
)

type GifRequest struct {
	storage.Gif
}

func (gr *GifRequest) Bind(r *http.Request) error {
	return nil
}
