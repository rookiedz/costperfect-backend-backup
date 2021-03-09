package api

import (
	"costperfect/backend/handlers/api"
	"net/http"

	"github.com/go-chi/chi"
)

//TestAPIRouter ...
func TestAPIRouter() http.Handler {
	var r *chi.Mux
	var hdlTest api.Test

	r = chi.NewRouter()
	hdlTest = api.NewTest()
	r.Route("/", func(r chi.Router) {
		r.Get("/", hdlTest.Get)
		r.Post("/", hdlTest.Create)
	})
	return r
}
