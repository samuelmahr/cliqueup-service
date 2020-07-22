package routers

import (
	"github.com/gorilla/mux"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"github.com/samuelmahr/cliqueup-service/internal/controllers"
	"github.com/samuelmahr/cliqueup-service/internal/repo"
)

type V1Router struct {
	config *configuration.AppConfig
	uRepo  repo.UsersRepoType
}

func NewV1Router(c *configuration.AppConfig, uRepo repo.UsersRepoType) V1Router {
	return V1Router{config: c, uRepo: uRepo}
}

//InitRoutes initialize all routes
func (v *V1Router) Register(root *mux.Router) {
	r := root.PathPrefix("/v1").Subrouter()

	usersController := controllers.NewV1UsersController(v.config, v.uRepo)
	usersController.RegisterRoutes(r)
}
