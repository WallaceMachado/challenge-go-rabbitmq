package interfaces

import "github.com/wallacemachado/challenge-go-rabbitmq/src/models"

type PersonRepository interface {
	CreatePerson(person *models.Person) (interface{}, error)
}
