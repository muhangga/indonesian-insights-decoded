package main

import (
	"fmt"
	"log"

	"github.com/muhangga/config"
	"github.com/muhangga/internal/app"
)

func main() {
	configuration := config.New(".env")

	initiate := app.NewInitializedServer(configuration)

	err := initiate.Run(fmt.Sprintf(":%s", configuration.Get("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
