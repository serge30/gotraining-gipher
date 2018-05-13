package web

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/serge30/gotraining-gipher/storage"
)

type GifResponse struct {
	storage.Gif
}

func (gr *GifResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type GifListResponse []*GifResponse

func NewGifListResponse(gifs []storage.Gif) []render.Renderer {
	list := []render.Renderer{}
	for _, gif := range gifs {
		list = append(list, &GifResponse{gif})
	}
	return list
}
