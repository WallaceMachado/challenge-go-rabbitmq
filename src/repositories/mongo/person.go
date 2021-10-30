package repositories

import (
	"context"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Publicacoes representa um repositório de publicações
type RepositoryPerson struct {
	Collection *mongo.Collection
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NewRepositoryPerson(c *mongo.Client) *RepositoryPerson {

	personCollection := c.Database("challenge-go-rabbitmq-db").Collection("person")
	return &RepositoryPerson{personCollection}
}

//var personCollection = database.Db().Database("challenge-go-rabbitmq-db").Collection("person")

func (r *RepositoryPerson) CreatePerson(person *models.Person) (interface{}, error) {

	insertResult, err := r.Collection.InsertOne(context.TODO(), person)
	if err != nil {
		return nil, err

	}

	return insertResult, nil
}

func (r *RepositoryPerson) GetPersonByName(name string) (*models.Person, error) {
	var person models.Person

	filter := bson.M{"name": name}

	if err := r.Collection.FindOne(context.TODO(), filter).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return &person, nil
		} else {
			return nil, err
		}

	}

	return &person, nil

}

func (r *RepositoryPerson) GetPersonById(id string) (*models.Person, error) {
	var person models.Person

	filter := bson.M{"_id": id}

	if err := r.Collection.FindOne(context.TODO(), filter).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return &person, nil
		} else {
			return nil, err
		}

	}

	return &person, nil

}

func (r *RepositoryPerson) ListAllPeople() (*[]models.Person, error) {
	var people []models.Person
	result, err := r.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {

		return nil, err

	}
	for result.Next(context.TODO()) {

		var person models.Person
		err := result.Decode(&person)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}
	result.Close(context.TODO())

	return &people, nil
}
