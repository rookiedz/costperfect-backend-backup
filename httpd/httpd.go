package httpd

import (
	"costperfect/config"
	"costperfect/routes"
	"costperfect/stores/mariadb"
	"log"
	"net/http"
	"time"
)

//Start ...
func Start() {
	var c config.Config
	var r http.Handler

	r = routes.InitialRouter()
	c = config.LoadConfiguration("config.json")
	mariadb.Connect(c.DataSourceName)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
