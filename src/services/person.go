package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/queue"
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
		return nil, errors.New("Person already exist")
	}

	result, err := p.repository.CreatePerson(person)
	if err != nil {
		return nil, err
	}

	messageCreatePersonToQueue := fmt.Sprintf("Cadastro da Pessoa com Id: %s", person.ID)

	data, _ := json.Marshal(messageCreatePersonToQueue)

	connection := queue.Connect()
	queue.Notify(data, config.RabbitmqExchangePerson, "", connection)

	return result, nil

}

func (p *PersonService) ListAllPeople() (*[]models.Person, error) {

	people, err := p.repository.ListAllPeople()
	if err != nil {
		return nil, err
	}
	messageUpdatePersonToQueue := fmt.Sprintf("Busca por todas Pessoas cadastradas")

	data, _ := json.Marshal(messageUpdatePersonToQueue)

	connection := queue.Connect()
	queue.Notify(data, config.RabbitmqExchangePerson, "", connection)

	return people, nil
}

func (p *PersonService) GetPersonById(id string) (*models.Person, error) {
	person, err := p.repository.GetPersonById(id)
	if err != nil {
		return nil, err
	}
	if person.ID == "" {
		return nil, errors.New("non-existent person")
	}

	messageGetPersonByIdToQueue := fmt.Sprintf("Busca da Pessoa com Id: %s", person.ID)

	data, _ := json.Marshal(messageGetPersonByIdToQueue)

	connection := queue.Connect()
	queue.Notify(data, config.RabbitmqExchangePerson, "", connection)

	return person, nil
}

func (p *PersonService) UpdatePerson(person *models.Person) error {
	personByID, err := p.repository.GetPersonById(person.ID)
	if err != nil {
		return err
	}
	if personByID.ID == "" {
		return errors.New("non-existent person")
	}

	personByName, err := p.repository.GetPersonByName(person.Name)
	if err != nil {
		return err
	}
	if personByName.ID != "" && personByName.ID != personByID.ID {
		return errors.New("there is already another person with that name")
	}

	err = p.repository.UpdatePerson(person)
	if err != nil {
		return err
	}

	messageUpdatePersonToQueue := fmt.Sprintf("Edição da Pessoa com Id: %s", person.ID)

	data, _ := json.Marshal(messageUpdatePersonToQueue)

	connection := queue.Connect()
	queue.Notify(data, config.RabbitmqExchangePerson, "", connection)

	return nil
}

func (p *PersonService) DeletePerson(id string) error {
	person, err := p.repository.GetPersonById(id)
	if err != nil {
		return err
	}
	if person.ID == "" {
		return errors.New("non-existent person")
	}

	err = p.repository.DeletePerson(id)
	if err != nil {
		return err
	}

	messageUpdatePersonToQueue := fmt.Sprintf("Exclusão da Pessoa com Id: %s", person.ID)

	data, _ := json.Marshal(messageUpdatePersonToQueue)

	connection := queue.Connect()
	queue.Notify(data, config.RabbitmqExchangePerson, "", connection)

	return nil
}
