package api

import (
	handlers "costperfect/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//UserAPI ...
func UserAPI() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetUsers)
		r.Post("/", handlers.CreateUser)
		r.Delete("/", handlers.DeleteUsers)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", handlers.GetUser)
		r.Put("/", handlers.UpdateUser)
		r.Delete("/", handlers.DeleteUser)
	})
	return r
}
