package httpd

import (
	"costperfect/backend/config"
	"costperfect/backend/routes"
	"costperfect/backend/stores/mariadb"
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
	mariadb.Connect(c.MariaDB.DataSourceName)

	s := &http.Server{
		Addr:           ":80",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(s.ListenAndServe())
}
