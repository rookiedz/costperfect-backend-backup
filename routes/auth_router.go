package routes

import (
	"costperfect/backend/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

//AuthRouter ...
func AuthRouter() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.Auth)
	})
	return r
}
