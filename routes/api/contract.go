package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//ContractAPIRouter ...
func ContractAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlContract api.Contract

	r = chi.NewRouter()
	hdlContract = api.NewContract()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlContract.All)
		r.Post("/", hdlContract.Create)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlContract.Get)
		r.Put("/", hdlContract.Update)
		r.Delete("/", hdlContract.Delete)
	})
	return r
}
