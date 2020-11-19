package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//OwnerAPIRouter ...
func OwnerAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlOwner api.Owner

	r = chi.NewRouter()
	hdlOwner = api.NewOwner()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlOwner.All)
		r.Post("/", hdlOwner.Create)
		r.Delete("/", hdlOwner.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlOwner.Get)
		r.Put("/", hdlOwner.Update)
		r.Delete("/", hdlOwner.Delete)
	})
	return r
}
