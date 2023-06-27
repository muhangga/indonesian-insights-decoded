package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/muhangga/api"
	"github.com/muhangga/db"
	"github.com/muhangga/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	producsRepo := repository.NewProductRepository(db)

	products := []repository.Product{
		{Name: "Product 1", Price: 10.000},
		{Name: "Product 2", Price: 20.000},
		{Name: "Product 3", Price: 20.400},
	}

	for _, product := range products {
		err := producsRepo.SaveProduct(product)
		if err != nil {
			log.Fatal(err)
		}

		product, err := producsRepo.FetchProduct()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(product)
	}

	mainAPI := api.NewAPI(*producsRepo)
	mainAPI.Start()
}
