package configuration

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Log                            log.Logger
	DatabaseURL                    string
	PostgresMaxOpenConns           int
	PostgresMaxIdleConns           int
	PostgresMaxConnLifetimeSeconds int

	TestDatabaseURL string
}

var globalConfig AppConfig

func Configuration() *AppConfig {
	return &globalConfig
}

func Configure() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxConns, err := strconv.Atoi(getEnv("POSTGRES_MAX_OPEN_CONNS", "10"))
	if err != nil {
		log.Fatal(err)
	}

	maxIdle, err := strconv.Atoi(getEnv("POSTGRES_MAX_IDLE_CONNS", "5"))
	if err != nil {
		log.Fatal(err)
	}

	maxLifetime, err := strconv.Atoi(getEnv("POSTGRES_MAX_CONN_LIFETIME_SECONDS", "3600"))
	if err != nil {
		log.Fatal(err)
	}

	c := AppConfig{}
	c.DatabaseURL = os.Getenv("DATABASE_URL")
	c.PostgresMaxOpenConns = maxConns
	c.PostgresMaxIdleConns = maxIdle
	c.PostgresMaxConnLifetimeSeconds = maxLifetime
	c.TestDatabaseURL = os.Getenv("TEST_DATABASE_URL")

	c.Log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	c.Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	c.Log.SetLevel(log.WarnLevel)

	return &globalConfig, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
