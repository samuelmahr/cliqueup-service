package application

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"github.com/samuelmahr/cliqueup-service/internal/repo"
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
	db, err := sqlx.Connect("postgres", c.DatabaseURL)
	if err != nil {
		log.Fatal("can't connect to db")
	}

	db.SetConnMaxLifetime(time.Duration(c.PostgresMaxConnLifetimeSeconds))
	db.SetMaxIdleConns(c.PostgresMaxIdleConns)
	db.SetMaxOpenConns(c.PostgresMaxOpenConns)

	userRepo := repo.NewUsersRepository(db)
	rootRouter := mux.NewRouter()
	r := routers.NewV1Router(c, userRepo)
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
