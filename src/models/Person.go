package models

import (
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty" valid:"notnull,uuid"`
	Name      string    `bson:"name" json:"name" valid:"notnull"`
	Gender    string    `bson:"gender" json:"gender" valid:"notnull"`
	Weight    float64   `bson:"weight" json:"weight" valid:"notnull,float"`
	Height    float64   `bson:"height" json:"height" valid:"notnull,float"`
	IMC       float64   `bson:"imc" json:"imc" valid:"notnull,float"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"createdAt" valid:"-"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updatedAt" valid:"-"`
}

func NewPerson(name string, gender string, weight float64, height float64, imc float64) (*Person, error) {
	person := &Person{
		Name:   name,
		Gender: gender,
		Weight: weight,
		Height: height,
		IMC:    imc,
	}

	person.ID = uuid.NewV4().String()
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()

	err := person.prepare()

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) prepare() error {

	p.Name = strings.TrimSpace(p.Name)
	p.Name = strings.ToLower(p.Name)
	p.Gender = strings.TrimSpace(p.Gender)
	p.Gender = strings.ToLower(p.Gender)

	err := p.validate()

	if err != nil {
		return err
	}

	return nil

}

func (p *Person) validate() error {

	govalidator.SetFieldsRequiredByDefault(true)

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}

func ValidatePersonID(id string) error {

	err := govalidator.IsUUID(id)

	if err == false {
		return errors.New("invalid ID")
	}

	return nil
}
