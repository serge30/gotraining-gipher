package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/serge30/gotraining-gipher/storage"
)

type GifRouter struct {
	storage storage.Storage
}

func GetRouters(storage storage.Storage) http.Handler {
	gr := GifRouter{storage}

	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/gifs", func(r chi.Router) {
		r.Get("/", gr.ListGifs)
		r.Post("/", gr.CreateGif) // POST /gifs

		r.Route("/{gifID}", func(r chi.Router) {
			r.Use(gr.GifCtx)            // Load the *Gif on the request context
			r.Get("/", gr.GetGif)       // GET /gifs/123
			r.Put("/", gr.UpdateGif)    // PUT /gifs/123
			r.Delete("/", gr.DeleteGif) // DELETE /gifs/123
		})
	})

	return r
}

func (gr *GifRouter) GifCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var gif storage.Gif
		var err error

		if gifIDStr := chi.URLParam(r, "gifID"); gifIDStr != "" {
			var gifID int
			gifID, err = strconv.Atoi(gifIDStr)
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
			gif, err = gr.storage.GetItem(gifID)
		}

		if err != nil {
			render.Render(w, r, ConverStorageError(err))
			return
		}

		ctx := context.WithValue(r.Context(), "gif", gif)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (gr *GifRouter) ListGifs(w http.ResponseWriter, r *http.Request) {
	gifs, err := gr.storage.GetItems()
	if err != nil {
		render.Render(w, r, ConverStorageError(err))
		return
	}

	if err := render.RenderList(w, r, NewGifListResponse(gifs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (gr *GifRouter) GetGif(w http.ResponseWriter, r *http.Request) {
	gif := r.Context().Value("gif").(storage.Gif)

	if err := render.Render(w, r, &GifResponse{gif}); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (gr *GifRouter) CreateGif(w http.ResponseWriter, r *http.Request) {
	data := &GifRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	gif := data.Gif

	newGif, err := gr.storage.CreateItem(gif)
	if err != nil {
		render.Render(w, r, ConverStorageError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &GifResponse{newGif})
}

func (gr *GifRouter) UpdateGif(w http.ResponseWriter, r *http.Request) {
	gif := r.Context().Value("gif").(storage.Gif)
	gifID := gif.ID

	data := &GifRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	gif = data.Gif

	newGif, err := gr.storage.UpdateItem(gifID, gif)
	if err != nil {
		render.Render(w, r, ConverStorageError(err))
		return
	}

	render.Render(w, r, &GifResponse{newGif})
}

func (gr *GifRouter) DeleteGif(w http.ResponseWriter, r *http.Request) {
	gif := r.Context().Value("gif").(storage.Gif)

	err := gr.storage.DeleteItem(gif.ID)
	if err != nil {
		render.Render(w, r, ConverStorageError(err))
		return
	}

	render.NoContent(w, r)
}
