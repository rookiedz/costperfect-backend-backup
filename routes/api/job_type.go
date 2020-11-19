package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//JobTypeAPIRouter ...
func JobTypeAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlJobType api.JobType

	r = chi.NewRouter()
	hdlJobType = api.NewJobType()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlJobType.All)
		r.Post("/", hdlJobType.Create)
		r.Delete("/", hdlJobType.DeleteByIDs)
	})
	r.Route("/{id:[0-9]+}", func(r chi.Router) {
		r.Get("/", hdlJobType.Get)
		r.Put("/", hdlJobType.Update)
		r.Delete("/", hdlJobType.Delete)
	})
	return r
}
