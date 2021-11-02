package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/database"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	repositories "github.com/wallacemachado/challenge-go-rabbitmq/src/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/satori/go.uuid"
)

var person = &models.Person{
	ID:        uuid.NewV4().String(),
	Name:      "teste",
	Weight:    80,
	Height:    175,
	IMC:       20,
	Gender:    "Masculino",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestCreatePerson(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		result, err := repository.CreatePerson(person)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("error duplicate name", func(t *testing.T) {
		person.ID = uuid.NewV4().String()

		result, err := repository.CreatePerson(person)

		assert.Nil(t, result)
		assert.True(t, mongo.IsDuplicateKeyError(err))

	})

	repository.Collection.Drop(context.TODO())

}

func TestGetPersonByName(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		repository.CreatePerson(person)

		result, err := repository.GetPersonByName(person.Name)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, result.ID, person.ID)

	})

	t.Run("error non-existent name", func(t *testing.T) {
		person.ID = uuid.NewV4().String()

		result, err := repository.GetPersonByName("non-existent name")

		assert.Empty(t, result)
		assert.Nil(t, err)

	})
	repository.Collection.Drop(context.TODO())
}

func TestGetPersonById(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		repository.CreatePerson(person)

		result, err := repository.GetPersonById(person.ID)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, result.ID, person.ID)

	})

	t.Run("error non-existent ID", func(t *testing.T) {
		person.ID = uuid.NewV4().String()

		result, err := repository.GetPersonById("non-existent ID")

		assert.Empty(t, result)
		assert.Nil(t, err)

	})
	repository.Collection.Drop(context.TODO())
}

func TestGetAllPeople(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		repository.CreatePerson(person)

		result, err := repository.ListAllPeople()

		assert.NotNil(t, result)
		assert.Nil(t, err)

	})

	repository.Collection.Drop(context.TODO())
}

func TestUpdatePerson(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		repository.CreatePerson(person)
		person.Name = "updated name"
		person.Weight = 81
		person.Height = 174
		person.IMC = 21
		person.Gender = "feminino"

		err := repository.UpdatePerson(person)

		assert.Nil(t, err)

	})

	t.Run("error duplicate name", func(t *testing.T) {
		person.ID = uuid.NewV4().String()
		person.Name = "teste"

		repository.CreatePerson(person)

		person.Name = "updated name"

		err := repository.UpdatePerson(person)

		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))

	})

	t.Run("error non-existent ID", func(t *testing.T) {

		err := repository.DeletePerson("non-existent ID")

		assert.NotNil(t, err)

	})

	repository.Collection.Drop(context.TODO())
}

func TestDeletePerson(t *testing.T) {
	config.Init()
	db := database.Db()
	defer database.DbClose()

	repository := repositories.NewRepositoryPerson(db, "challenge-go-rabbitmq-dbTeste")

	t.Run("Success", func(t *testing.T) {

		repository.CreatePerson(person)

		err := repository.DeletePerson(person.ID)

		assert.Nil(t, err)

	})

	t.Run("error non-existent ID", func(t *testing.T) {

		err := repository.DeletePerson("non-existent ID")

		assert.NotNil(t, err)

	})
	repository.Collection.Drop(context.TODO())
}
