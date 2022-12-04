package routers

import (
	"github.com/gorilla/mux"
)

type MyRouter mux.Router

func (r *MyRouter) Routes(route *mux.Router) *mux.Router {
	route.HandleFunc("/users", GetUsers).Methods("GET")

	return route
}
