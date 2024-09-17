package main

import (
	"ciyana/controller"
	"ciyana/delivery/route"
	"ciyana/repository"
	"ciyana/usecase"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database("Ciyana")
	itemCollection := db.Collection("items")
	categoryCollection := db.Collection("categories")

	itemRepo := repository.NewItemRepository(itemCollection, context.TODO())
	categoryRepo := repository.NewCategoryRepository(categoryCollection, context.TODO())

	itemUsecase := usecase.NewItemUsecase(itemRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	controller := controller.NewController(itemUsecase, categoryUsecase)

	route := route.SetupRouter(controller)
	route.Run("localhost:3000")

}
