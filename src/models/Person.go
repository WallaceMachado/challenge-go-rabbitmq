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
	Weight    float32   `bson:"weight" json:"weight" valid:"notnull,float"`
	Height    float32   `bson:"height" json:"height" valid:"notnull,float"`
	IMC       float32   `bson:"imc" json:"imc" valid:"notnull,float"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"createdAt" valid:"-"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updatedAt" valid:"-"`
}

func NewPerson(p *Person, stage string) (*Person, error) {
	person := &Person{
		ID:     p.ID,
		Name:   p.Name,
		Gender: p.Gender,
		Weight: p.Weight,
		Height: p.Height,
		IMC:    p.IMC,
	}

	if stage == "create" {
		person.ID = uuid.NewV4().String()
		person.CreatedAt = time.Now()
		person.UpdatedAt = time.Now()
	} else if stage == "update" {
		person.UpdatedAt = time.Now()
	}

	if err := person.prepare(); err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) prepare() error {

	p.Name = strings.TrimSpace(p.Name)
	p.Name = strings.ToLower(p.Name)
	p.Gender = strings.TrimSpace(p.Gender)
	p.Gender = strings.ToLower(p.Gender)

	if err := p.validate(); err != nil {
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

	if err := validatePersonName(p.Name); err != nil {
		return err
	}

	if err := validatePersonGender(p.Gender); err != nil {
		return err
	}

	if err := validatePersonWeight(p.Weight); err != nil {
		return err
	}

	if err := validatePersonHeight(p.Height); err != nil {
		return err
	}

	if err := validatePersonIMC(p.IMC); err != nil {
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

func validatePersonGender(gender string) error {

	genderTypes := []string{"male", "female"}
	validateGender := false
	for _, v := range genderTypes {
		if gender == v {
			validateGender = true
		}
	}

	if !validateGender {
		return errors.New(`invalid gender, please use "male" or "female"`)
	}

	return nil
}

func validatePersonName(name string) error {

	if len(name) < 3 || len(name) > 100 {

		return errors.New("The name must be between 3 and 100 characters.")
	}

	return nil
}
func validatePersonHeight(height float32) error {

	if height < 30 || height > 250 {
		return errors.New("The height must be informed in centimeters and must be between 30cm and 250cm")
	}

	return nil
}

func validatePersonWeight(weight float32) error {

	if weight < 10 || weight > 900 {
		return errors.New("The weight must be informed in kg and must be between 10kg and 900kg")
	}

	return nil
}

func validatePersonIMC(imc float32) error {

	if imc < 10 || imc > 100 {
		return errors.New("The IMC must be between 10 and 100")
	}

	return nil
}
