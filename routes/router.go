package routes

import (
	"costperfect/backend/routes/api"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

//InitialRouter ...
func InitialRouter() http.Handler {
	var r *chi.Mux
	r = chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	api.Initial()
	r.Route("/api", func(r chi.Router) {
		r.Mount("/employers", api.EmployerAPIRouter())
		r.Mount("/contracts", api.ContractAPIRouter())
		r.Mount("/contractors", api.ContractorAPIRouter())
		r.Mount("/jobs/groups", api.JobGroupAPIRouter())
		r.Mount("/jobs/types", api.JobTypeAPIRouter())
		r.Mount("/jobs", api.JobAPIRouter())
		r.Mount("/installments", api.InstallmentAPIRouter())
		r.Mount("/owners", api.OwnerAPIRouter())
		r.Mount("/projects", api.ProjectAPIRouter())
		r.Mount("/users", api.UserAPIRouter())
		//r.Mount("/authors", api.AuthorAPI())
	})
	r.Mount("/auth", AuthRouter())

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "html"))
	fmt.Println(filesDir)
	//FileServer(r, "/", filesDir)

	return r
}

//FileServer ...
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
