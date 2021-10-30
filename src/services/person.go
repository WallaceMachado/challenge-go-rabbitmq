package services

import (
	"errors"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/repositories/interfaces"
)

type PersonService struct {
	repository interfaces.PersonRepository
}

func NewPersonService(repo interfaces.PersonRepository) *PersonService {
	return &PersonService{
		repository: repo,
	}
}

func (p *PersonService) CreatePerson(person *models.Person) (interface{}, error) {

	personByName, err := p.repository.GetPersonByName(person.Name)
	if err != nil {
		return nil, err
	}
	if personByName.ID != "" {
		return nil, errors.New("Person already exists")
	}

	return p.repository.CreatePerson(person)
}
