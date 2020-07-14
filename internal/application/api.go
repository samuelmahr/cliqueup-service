package application

import (
	"github.com/gorilla/mux"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"github.com/samuelmahr/cliqueup-service/internal/routers"
	"log"
	"net/http"
	"time"
)

type APIApplication struct {
	config *configuration.AppConfig
	srv    *http.Server
}

func NewAPIApplication(c *configuration.AppConfig) *APIApplication {
	rootRouter := mux.NewRouter()
	r := routers.NewV1Router(c)
	r.Register(rootRouter)

	srv := &http.Server{
		Handler: rootRouter,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &APIApplication{
		config: c,
		srv:    srv,
	}
}

func (a *APIApplication) Run() {
	log.Fatal(a.srv.ListenAndServe())
}
