package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	"github.com/samuelmahr/cliqueup-service/internal/models"
	"github.com/samuelmahr/cliqueup-service/internal/repo"
	"net/http"
)

type errorTyper interface {
	ErrorType() string
}

type junk struct {
	Hello string `json:"hello"`
}

type V1UsersController struct {
	config *configuration.AppConfig
	repo   repo.UsersRepoType
}

func NewV1UsersController(c *configuration.AppConfig, uRepo repo.UsersRepoType) V1UsersController {
	return V1UsersController{
		config: c,
		repo:   uRepo,
	}
}

func (a *V1UsersController) RegisterRoutes(v1 *mux.Router) {
	v1.Path("/users").Name("CreateUser").Handler(http.HandlerFunc(a.CreateUser)).Methods(http.MethodPost)
}

func (a *V1UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	newUser := models.UsersCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondError(ctx, w, http.StatusBadRequest, "bad request", err)
		return
	}

	user, err := a.repo.CreateUser(ctx, newUser)
	if err != nil {
		respondError(ctx, w, http.StatusInternalServerError, "lmfao something happened", err)
		return
	}

	respondModel(ctx, w, http.StatusCreated, user)
	return
}
