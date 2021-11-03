package repositories

import (
	"context"
	"errors"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryPerson struct {
	Collection *mongo.Collection
}

func NewRepositoryPerson(c *mongo.Client, dbName string) *RepositoryPerson {
	newIndex := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	}
	personCollection := c.Database(dbName).Collection(config.DbCollection)
	personCollection.Indexes().CreateOne(context.TODO(), newIndex)
	return &RepositoryPerson{personCollection}
}

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

func (r *RepositoryPerson) ListAllPeople() ([]*models.Person, error) {
	var people []*models.Person
	result, err := r.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {

		return nil, err

	}
	for result.Next(context.TODO()) {

		var person *models.Person
		err := result.Decode(&person)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}
	result.Close(context.TODO())

	return people, nil
}

func (r *RepositoryPerson) UpdatePerson(person *models.Person) error {

	filter := bson.M{"_id": person.ID}

	update := bson.M{"$set": bson.M{
		"name":       person.Name,
		"gender":     person.Gender,
		"heigth":     person.Height,
		"weigth":     person.Weight,
		"imc":        person.IMC,
		"updated_at": person.UpdatedAt,
	}}

	result, err := r.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("invalid ID")
	}

	return nil
}

func (r *RepositoryPerson) DeletePerson(id string) error {

	filter := bson.M{"_id": id}

	result, err := r.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("invalid ID")
	}

	return nil
}
