package routes

import (
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/controllers"
)

var routesPerson = []Router{
	{
		URI:    "/person",
		Metodo: http.MethodGet,
		Funcao: controllers.GetAllPeople,
	},

	{
		URI:    "/person",
		Metodo: http.MethodPost,
		Funcao: controllers.CreatePerson,
	},
}
