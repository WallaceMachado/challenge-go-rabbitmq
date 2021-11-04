package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
)

func TestNewAccount(t *testing.T) {

	t.Run("NewPerson with ID: success", func(t *testing.T) {
		p := &models.Person{
			ID:     "c4ed67f0-93fb-4c50-aacd-30e9b046b725",
			Name:   "teste",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}

		result, err := models.NewPerson(p, "update")
		require.Nil(t, err)
		assert.NotNil(t, result.UpdatedAt)
		assert.Equal(t, result.ID, p.ID)
	})

	t.Run("NewPerson without ID: success", func(t *testing.T) {

		p := &models.Person{
			Name:   "teste",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		require.Nil(t, err)
		assert.NotNil(t, result.CreatedAt)
		assert.NotNil(t, result.UpdatedAt)
		assert.NotEmpty(t, result.ID)

	})

	t.Run("Incorrect ID: id cannot be null", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}
		err := models.ValidatePersonID(p.ID)
		assert.EqualError(t, err, "invalid ID")

	})

	t.Run("Incorrect Name: name cannot be less than 3 characters", func(t *testing.T) {
		p := &models.Person{
			Name:   "  ab ",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}

		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The name must be between 3 and 100 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Name: name cannot be longer than 100 characters", func(t *testing.T) {
		p := &models.Person{
			Name:   "",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}
		p.Name = generateString(101)
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The name must be between 3 and 100 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect gender: gender must be male or female", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 80,
			Height: 175,
			IMC:    20,
			Gender: "male",
		}
		p.Gender = "incorrect gender"
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, `invalid gender, please use "male" or "female"`)
		assert.Nil(t, result)

	})

	t.Run("Incorrect Height: height cannot be less than 30cm", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 80,
			Height: 29,
			IMC:    20,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The height must be informed in centimeters and must be between 30cm and 250cm")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Height: height cannot be longer than 250 characters", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 80,
			Height: 251,
			IMC:    20,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The height must be informed in centimeters and must be between 30cm and 250cm")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Weight: weight cannot be less than 10kg", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 9,
			Height: 30,
			IMC:    20,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The weight must be informed in kg and must be between 10kg and 900kg")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Weight: Weight cannot be longer than 900kg", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 901,
			Height: 250,
			IMC:    20,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The weight must be informed in kg and must be between 10kg and 900kg")
		assert.Nil(t, result)

	})

	t.Run("Incorrect IMC: IMC cannot be less than 10", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 30,
			Height: 30,
			IMC:    9,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The IMC must be between 10 and 100")
		assert.Nil(t, result)

	})

	t.Run("Incorrect IMC: IMC cannot be longer than 100", func(t *testing.T) {
		p := &models.Person{
			Name:   "teste",
			Weight: 900,
			Height: 250,
			IMC:    101,
			Gender: "male",
		}
		result, err := models.NewPerson(p, "create")
		assert.EqualError(t, err, "The IMC must be between 10 and 100")
		assert.Nil(t, result)

	})

}

func generateString(n int) string {
	a := "a"
	s := ""
	for i := 0; i < n; i++ {
		s = s + a
	}

	return s
}
