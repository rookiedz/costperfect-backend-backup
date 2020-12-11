package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//EmployerAPIRouter ...
func EmployerAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlEmployer api.Employer

	r = chi.NewRouter()
	hdlEmployer = api.NewEmployer()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlEmployer.All)
		r.Post("/", hdlEmployer.Create)
		r.Delete("/", hdlEmployer.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlEmployer.Get)
		r.Put("/", hdlEmployer.Update)
		r.Delete("/", hdlEmployer.Delete)
	})
	return r
}
