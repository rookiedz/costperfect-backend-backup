package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//UserAPI ...
func UserAPI() http.Handler {
	var r *chi.Mux
	var hdlUser api.User

	r = chi.NewRouter()
	hdlUser = api.NewUser()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlUser.All)
		r.Post("/", hdlUser.Create)
		r.Delete("/", hdlUser.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlUser.Get)
		r.Put("/", hdlUser.Update)
		r.Delete("/", hdlUser.Delete)
	})
	return r
}
