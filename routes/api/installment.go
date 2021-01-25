package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//InstallmentAPIRouter ...
func InstallmentAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlInstallment api.Installment

	r = chi.NewRouter()
	hdlInstallment = api.NewInstallment()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlInstallment.All)
		r.Post("/", hdlInstallment.Create)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlInstallment.Get)
		r.Put("/", hdlInstallment.Update)
		r.Delete("/", hdlInstallment.Delete)
	})
	return r
}
