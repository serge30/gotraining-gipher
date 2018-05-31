package web

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/serge30/gotraining-gipher/storage"
)

// GifResponse is the response payload for Gif data model.
type GifResponse struct {
	storage.Gif
}

// Render is standard methos to implement render.Render interface.
func (gr *GifResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// GifListResponse is the response payload for list of Gif data models.
type GifListResponse []*GifResponse

// NewGifListResponse creates new slice of renderers for slice of Gifs.
func NewGifListResponse(gifs []storage.Gif) []render.Renderer {
	list := []render.Renderer{}
	for _, gif := range gifs {
		list = append(list, &GifResponse{gif})
	}
	return list
}
