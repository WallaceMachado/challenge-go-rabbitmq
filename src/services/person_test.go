package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/services"
)

type RepositoryPersonInMemory struct {
	repo []*models.Person
}

func (r *RepositoryPersonInMemory) CreatePerson(person *models.Person) (interface{}, error) {

	r.repo = append(r.repo, person)

	return r.repo[len(r.repo)-1].ID, nil
}

func (r *RepositoryPersonInMemory) GetPersonByName(name string) (*models.Person, error) {
	var person *models.Person

	for _, v := range r.repo {
		if v.Name == name {
			person = v
		}
	}
	if person == nil {
		return &models.Person{}, nil
	}

	return person, nil

}

func (r *RepositoryPersonInMemory) GetPersonById(id string) (*models.Person, error) {
	var person *models.Person

	for _, v := range r.repo {
		if v.ID == id {
			person = v
		}
	}

	if person == nil {
		return &models.Person{}, nil
	}

	return person, nil

}

func (r *RepositoryPersonInMemory) ListAllPeople() ([]*models.Person, error) {

	return r.repo, nil
}

func (r *RepositoryPersonInMemory) UpdatePerson(person *models.Person) error {

	for i, v := range r.repo {
		if v.ID == person.ID {
			r.repo[i] = v
		}
	}

	return nil
}

func (r *RepositoryPersonInMemory) DeletePerson(id string) error {

	for i, v := range r.repo {
		if v.ID == id {
			r.repo = append(r.repo[:i], r.repo[i+1:]...)
		}
	}

	return nil
}

var p = &models.Person{
	ID:        "c4ed67f0-93fb-4c50-aacd-30e9b046b725",
	Name:      "teste",
	Weight:    80,
	Height:    175,
	IMC:       20,
	Gender:    "Masculino",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var p2 = &models.Person{
	ID:        "c4ed67f0-93fb-4c50-aacd-30e9b046b733",
	Name:      "teste2",
	Weight:    82,
	Height:    185,
	IMC:       22,
	Gender:    "Feminino",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestCreatePerson(t *testing.T) {
	config.Init()

	var repo []*models.Person

	service := services.NewPersonService(&RepositoryPersonInMemory{repo})

	t.Run("Success", func(t *testing.T) {

		result, err := service.CreatePerson(p)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("error duplicate name", func(t *testing.T) {
		p2.Name = p.Name

		result, err := service.CreatePerson(p2)

		assert.Nil(t, result)
		assert.EqualError(t, err, "Person already exist")

	})

}

func TestGetPersonByid(t *testing.T) {
	config.Init()

	var repo []*models.Person

	service := services.NewPersonService(&RepositoryPersonInMemory{repo})

	t.Run("Success", func(t *testing.T) {

		service.CreatePerson(p)

		result, err := service.GetPersonById(p.ID)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, result.ID, p.ID)

	})

	t.Run("error non-existent name", func(t *testing.T) {

		result, err := service.GetPersonById("non-existent ID")

		assert.Nil(t, result)
		assert.EqualError(t, err, "non-existent person")

	})

}

func TestGetAllPeople(t *testing.T) {
	config.Init()

	var repo []*models.Person

	service := services.NewPersonService(&RepositoryPersonInMemory{repo})

	t.Run("Success", func(t *testing.T) {

		service.CreatePerson(p)

		result, err := service.ListAllPeople()

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Len(t, result, 1)

	})

}

func TestUpdatePerson(t *testing.T) {
	config.Init()

	var repo []*models.Person

	service := services.NewPersonService(&RepositoryPersonInMemory{repo})

	var p3 = &models.Person{
		ID:        "c4ed67f0-93fb-4c50-aacd-30e9b046b733",
		Name:      "updated name",
		Weight:    82,
		Height:    185,
		IMC:       22,
		Gender:    "Feminino",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {

		service.CreatePerson(p)

		p.Name = "updated name"
		p.Weight = 81
		p.Height = 174
		p.IMC = 21
		p.Gender = "feminino"

		err := service.UpdatePerson(p)

		assert.Nil(t, err)

	})

	t.Run("error non-existent ID", func(t *testing.T) {
		service.CreatePerson(p2)

		err := service.UpdatePerson(p3)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "there is already another person with that name")

	})

	t.Run("error non-existent ID", func(t *testing.T) {
		p3.ID = "non-existent ID"
		err := service.UpdatePerson(p3)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "non-existent person")

	})

}

func TestDeletePerson(t *testing.T) {
	config.Init()

	var repo []*models.Person

	service := services.NewPersonService(&RepositoryPersonInMemory{repo})

	t.Run("Success", func(t *testing.T) {

		service.CreatePerson(p)

		err := service.DeletePerson(p.ID)

		assert.Nil(t, err)

	})

	t.Run("error non-existent ID", func(t *testing.T) {

		err := service.DeletePerson(p2.ID)

		assert.NotNil(t, err)

	})

}
