package repository

import (
	"context"
	"log"

	"github.com/zdos/dodo_pizza/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PizzaDb interface {
	GetPizzaList(ctx context.Context) ([]domain.Pizza, error)
}

type pizzaRepo struct {
	pizzaCollection *mongo.Collection
}

func NewPizzaRepo(db *mongo.Database) *pizzaRepo {
	return &pizzaRepo{
		pizzaCollection: db.Collection("pizza"),
	}
}

func (r *pizzaRepo) GetPizzaList(ctx context.Context) ([]domain.Pizza, error) {
	pizzaList, err := r.pizzaCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	result := make([]domain.Pizza, 0)
	for pizzaList.Next(ctx) {
		pizza := new(domain.Pizza)
		if err := pizzaList.Decode(pizza); err != nil {
			log.Printf("decode err: %s. skip...", err.Error())
			continue
		}
		result = append(result, *pizza)
	}

	return result, nil
}
