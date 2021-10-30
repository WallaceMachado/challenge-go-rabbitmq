package services

import (
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

	return p.repository.CreatePerson(person)
}
