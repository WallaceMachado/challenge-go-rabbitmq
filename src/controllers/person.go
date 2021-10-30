package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/controllers/responses"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/database"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	repositories "github.com/wallacemachado/challenge-go-rabbitmq/src/repositories/mongo"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/services"
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

	var personCollection = database.Db().Database("challenge-go-rabbitmq-db").Collection("person")
	defer personCollection.Database().Client().Disconnect(context.TODO())

	repository := repositories.NewRepositoryPerson(personCollection)

	service := services.NewPersonService(repository)

	result, err := service.CreatePerson(newPerson)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, result)

}
