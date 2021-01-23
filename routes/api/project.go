package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//ProjectAPIRouter ...
func ProjectAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlProject api.Project

	r = chi.NewRouter()
	hdlProject = api.NewProject()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlProject.All)
		r.Post("/", hdlProject.Create)
		r.Delete("/", hdlProject.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlProject.Get)
		r.Get("/contracts", hdlProject.Contracts)
		r.Get("/contractors", hdlProject.Contractors)
		r.Get("/employers", hdlProject.Employers)
		r.Put("/", hdlProject.Update)
		r.Delete("/", hdlProject.Delete)
	})
	return r
}
