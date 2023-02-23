package repository

import (
	"context"
	"log"
	"strings"

	"github.com/zdos/dodo_pizza/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PizzaDb interface {
	GetPizzaList(ctx context.Context, filter *domain.PizzaFitlerReq) ([]domain.Pizza, error)
}

type pizzaRepo struct {
	pizzaCollection *mongo.Collection
}

func NewPizzaRepo(db *mongo.Database) *pizzaRepo {
	return &pizzaRepo{
		pizzaCollection: db.Collection("pizza"),
	}
}

func (r *pizzaRepo) GetPizzaList(ctx context.Context, filter *domain.PizzaFitlerReq) ([]domain.Pizza, error) {
	mfilter := bson.D{}
	orderByValue := 1
	if strings.ToLower(*filter.OrderBy) == "desc" {
		orderByValue = -1
	}

	if filter.Category != nil {
		mfilter = append(mfilter, primitive.E{Key: "category", Value: *filter.Category})
	}

	opts := options.Find().SetSort(bson.D{{Key: *filter.SortBy, Value: orderByValue}})
	pizzaList, err := r.pizzaCollection.Find(ctx, mfilter, opts)
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
