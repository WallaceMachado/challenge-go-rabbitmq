package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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

	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db)

	service := services.NewPersonService(repository)

	result, err := service.CreatePerson(newPerson)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, result)

}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db)

	service := services.NewPersonService(repository)
	people, err := service.ListAllPeople()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, people)
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	if err := models.ValidatePersonID(id); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db)

	service := services.NewPersonService(repository)
	result, err := service.GetPersonById(id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}
