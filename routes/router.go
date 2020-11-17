package routes

import (
	"costperfect/backend/routes/api"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

//InitialRouter ...
func InitialRouter() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", api.UserAPI())
		r.Mount("/authors", api.AuthorAPI())
	})
	r.Mount("/auth", AuthRouter())
	return r
}
