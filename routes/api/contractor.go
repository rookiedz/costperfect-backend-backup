package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//ContractorAPIRouter ...
func ContractorAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlContractor api.Contractor

	r = chi.NewRouter()
	hdlContractor = api.NewContractor()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlContractor.All)
		r.Post("/", hdlContractor.Create)
		r.Delete("/", hdlContractor.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlContractor.Get)
		r.Put("/", hdlContractor.Update)
		r.Delete("/", hdlContractor.Delete)
	})
	return r
}
