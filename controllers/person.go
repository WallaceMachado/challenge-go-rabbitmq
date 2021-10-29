package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/controllers/responses"
	"github.com/wallacemachado/challenge-go-rabbitmq/models"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var person models.Person

	if err = json.Unmarshal(bodyRequest, &person); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	newPerson, err := models.NewPerson(person.Name, person.Gender, person.Weight, person.Height, person.IMC)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, newPerson)

}
