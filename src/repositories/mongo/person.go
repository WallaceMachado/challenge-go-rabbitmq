package repositories

import (
	"context"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Publicacoes representa um repositório de publicações
type RepositoryPerson struct {
	Collection *mongo.Collection
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NewRepositoryPerson(c *mongo.Collection) *RepositoryPerson {
	return &RepositoryPerson{c}
}

//var personCollection = database.Db().Database("challenge-go-rabbitmq-db").Collection("person")

func (r *RepositoryPerson) CreatePerson(person *models.Person) (interface{}, error) {

	insertResult, err := r.Collection.InsertOne(context.TODO(), person)
	if err != nil {
		return nil, err

	}

	return insertResult, nil
}
