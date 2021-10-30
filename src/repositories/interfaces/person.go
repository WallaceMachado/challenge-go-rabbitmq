package interfaces

import "github.com/wallacemachado/challenge-go-rabbitmq/src/models"

type PersonRepository interface {
	CreatePerson(person *models.Person) (interface{}, error)
	GetPersonById(id string) (*models.Person, error)
	GetPersonByName(name string) (*models.Person, error)
	ListAllPeople() (*[]models.Person, error)
	UpdatePerson(person *models.Person) error
}
