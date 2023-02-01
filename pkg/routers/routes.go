package routers

import (
	"database/sql"
	"net/http"

	"go-demoproject/pkg/controllers"

	"github.com/gorilla/mux"
)

type MyRouter mux.Router

func (r *MyRouter) Routes(route *mux.Router, db *sql.DB) *mux.Router {

	userController := controllers.NewUserController(db)
	route.HandleFunc("/users", userController.GetUsers()).Methods(http.MethodGet)
	route.HandleFunc("/user/create", userController.Create()).Methods(http.MethodPost)

	return route
}
