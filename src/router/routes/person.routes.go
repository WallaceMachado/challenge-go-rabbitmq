package routes

import (
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/controllers"
)

var routesPerson = []Router{

	{
		URI:    "/person",
		Metodo: http.MethodPost,
		Funcao: controllers.CreatePerson,
	},
	{
		URI:    "/person",
		Metodo: http.MethodGet,
		Funcao: controllers.GetAllPeople,
	},
	{
		URI:    "/person/{id}",
		Metodo: http.MethodGet,
		Funcao: controllers.GetPersonById,
	},
	{
		URI:    "/person/{id}",
		Metodo: http.MethodPut,
		Funcao: controllers.UpdatePerson,
	},
	{
		URI:    "/person/{id}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletePerson,
	},
}
