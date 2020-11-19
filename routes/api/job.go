package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//JobAPIRouter ...
func JobAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlJob api.Job

	r = chi.NewRouter()
	hdlJob = api.NewJob()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlJob.All)
		r.Post("/", hdlJob.Create)
		r.Delete("/", hdlJob.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlJob.Get)
		r.Put("/", hdlJob.Update)
		r.Delete("/", hdlJob.Delete)
	})
	return r
}
