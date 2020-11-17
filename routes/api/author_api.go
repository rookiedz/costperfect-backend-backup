package api

import (
	handlers "costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//AuthorAPI ...
func AuthorAPI() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetAuthors)
		r.Post("/", handlers.CreateAuthor)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", handlers.GetAuthor)
		r.Put("/", handlers.UpdateAuthor)
		r.Delete("/", handlers.DeleteAuthor)
	})
	return r
}
