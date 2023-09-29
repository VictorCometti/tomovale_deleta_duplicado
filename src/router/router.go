package router

import (
	"github.com/gorilla/mux"
	"tomovale_deleta_duplicado/src/router/routers"
)

func GetRouter() (r *mux.Router) {
	r = mux.NewRouter()
	routers.ConfigureRoute(r)
	return
}
