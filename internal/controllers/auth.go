package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/samuelmahr/cliqueup-service/internal/configuration"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type errorTyper interface {
	ErrorType() string
}

type junk struct {
	Hello string `json:"hello"`
}

type V1AuthController struct {
	config *configuration.AppConfig
}

func NewV1AuthController(c *configuration.AppConfig) V1AuthController {
	return V1AuthController{
		config: c,
	}
}

func (a *V1AuthController) RegisterRoutes(v1 *mux.Router) {
	v1.Path("/auth/createUser").Name("CreateUser").Handler(http.HandlerFunc(a.CreateUser)).Methods(http.MethodPost)
}

func (a *V1AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b, err := json.Marshal(junk{Hello: "world"})
	if err != nil {
		a.respondError(ctx, w, 500, "error generating response", err)
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write(b)
}

func (a *V1AuthController) respondError(ctx context.Context, w http.ResponseWriter, status int, message string, causer error) {
	resp := map[string]interface{}{
		"error": message,
	}

	if status >= 500 {
		a.config.Log.WithFields(log.Fields{
			"message": message,
			"causer":  causer,
		},
		).Error("yeeaah buddy")
	}

	if typer, ok := causer.(errorTyper); ok {
		resp["type"] = typer.ErrorType()
	}

	if errors.Cause(causer) == sql.ErrNoRows {
		// smahr 6/2 need to rethink below line so we are not returning "sql: no rows in result set" in the error message in response
		// resp["error"] = "not found"
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(status)
	}

	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}
