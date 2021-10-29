package routes

import (
	"fmt"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/controllers"
)

var routesPerson = []Router{
	{
		URI:    "/person",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Get Person")
		},
	},

	{
		URI:    "/person",
		Metodo: http.MethodPost,
		Funcao: controllers.CreatePerson,
	},
}
