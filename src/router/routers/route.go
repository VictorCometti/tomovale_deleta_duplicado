package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func ConfigureRoute(r *mux.Router) (router *mux.Router) {
	routes := tomovaleRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return

}
