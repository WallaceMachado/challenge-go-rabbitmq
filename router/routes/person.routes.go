package routes

import (
	"fmt"
	"net/http"
)

var routesPerson = []Router{
	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Get Person")
		},
	},
}
