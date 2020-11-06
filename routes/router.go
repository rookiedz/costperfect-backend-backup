package routes

import (
	"costperfect/routes/api"
	"net/http"

	"github.com/go-chi/chi"
)

//InitialRouter ...
func InitialRouter() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", api.UserAPI())
		r.Mount("/authors", api.AuthorAPI())
	})
	r.Mount("/auth", AuthRouter())
	return r
}
