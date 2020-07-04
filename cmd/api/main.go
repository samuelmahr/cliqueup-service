package main

import (
	"github.com/samuelmahr/cliqueup-service/internal/application"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"log"
)

func main() {
	c, err := configuration.Configure()
	if err != nil {
		log.Fatal(err)
	}

	app := application.NewAPIApplication(c)
	app.Run()
}
