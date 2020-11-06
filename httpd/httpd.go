package httpd

import (
	"costperfect/config"
	"costperfect/routes"
	"costperfect/stores/mariadb"
	"net/http"
)

//Start ...
func Start() {
	var c config.Config
	var r http.Handler

	r = routes.InitialRouter()
	c = config.LoadConfiguration("config.json")
	mariadb.Connect(c.DataSourceName)

	http.ListenAndServe(":8080", r)
}
