package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func respondError(ctx context.Context, w http.ResponseWriter, status int, message string, causer error) {
	resp := map[string]interface{}{
		"error": message,
	}

	if status >= 400 {
		log.WithFields(log.Fields{
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

func respondModel(ctx context.Context, w http.ResponseWriter, status int, model interface{}) {
	b, err := json.Marshal(model)
	if err != nil {
		respondError(ctx, w, 500, "error generating response", err)
	}

	w.WriteHeader(status)
	_, _ = w.Write(b)
}
