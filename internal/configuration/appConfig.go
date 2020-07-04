package configuration

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type AppConfig struct {
	Log log.Logger
}

var globalConfig AppConfig

func Configuration() *AppConfig {
	return &globalConfig
}

func Configure() (*AppConfig, error) {
	c := AppConfig{}

	c.Log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	c.Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	c.Log.SetLevel(log.WarnLevel)

	return &globalConfig, nil
}
