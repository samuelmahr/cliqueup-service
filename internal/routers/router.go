package routers

import (
	"github.com/gorilla/mux"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"github.com/samuelmahr/cliqueup-service/internal/controllers"
)

type V1Router struct {
	config *configuration.AppConfig
}

func NewV1Router(c *configuration.AppConfig) V1Router {
	return V1Router{config: c}
}

//InitRoutes initialize all routes
func (v *V1Router) Register(root *mux.Router) {
	r := root.PathPrefix("/v1").Subrouter()

	usersController := controllers.NewV1UsersController(v.config)
	usersController.RegisterRoutes(r)
}
