package main

import (
	"context"
	"log"
	"net/http"

	"github.com/zdos/dodo_pizza/internal/controller"
	"github.com/zdos/dodo_pizza/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	opts := options.Client().ApplyURI("mongodb+srv://admin:root@cluster0.bu48yln.mongodb.net/pizza_db?retryWrites=true&w=majority")

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("got error mongo: %s", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("got error ping mongo: %s", err.Error())
	}

	mongoPizzaDb := client.Database("pizza_db")

	pizzaRepo := repository.NewPizzaRepo(mongoPizzaDb)

	router := controller.NewRouter(pizzaRepo)

	httpSrv := http.Server{
		Addr:    "localhost:8089",
		Handler: router.Init(),
	}
	log.Printf("http server starting...")

	if err := httpSrv.ListenAndServe(); err != nil {
		log.Fatalf("got error while listenAndServe: %s", err.Error())
	}
}
