package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//JobGroupAPIRouter ...
func JobGroupAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlJobGroup api.JobGroup

	r = chi.NewRouter()
	hdlJobGroup = api.NewJobGroup()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlJobGroup.All)
		r.Post("/", hdlJobGroup.Create)
		r.Delete("/", hdlJobGroup.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlJobGroup.Get)
		r.Put("/", hdlJobGroup.Update)
		r.Delete("/", hdlJobGroup.Delete)
	})
	return r
}
