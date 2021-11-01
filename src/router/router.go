package router

import (
	"github.com/gorilla/mux"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/router/routes"
)

func Generate() *mux.Router {

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	return routes.SetUp(s)
}
