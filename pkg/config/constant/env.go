package config

import (
	"os"

	"github.com/fernandojosemoran/go-templates/pkg/helpers"
	"github.com/fernandojosemoran/go-templates/pkg/logger"
	"github.com/joho/godotenv"
)

// https://data.iana.org/time-zones/tzdb-2021a/zone1970.tab

type env struct {
	POSTGRES_DB_NAME  string
	POSTGRES_DB_TABLE string
	POSTGRES_DNS      string
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	PORT              int
	DEBUG             bool
}

func Env() *env {
	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}

	POSTGRES_DB_NAME_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_DB_NAME"),
		VariableName: "POSTGRES_DB_NAME",
	}

	POSTGRES_DB_TABLE_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_DB_TABLE"),
		VariableName: "POSTGRES_DB_TABLE",
	}

	POSTGRES_DNS_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_DNS"),
		VariableName: "POSTGRES_DNS",
	}

	POSTGRES_HOST_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_HOST"),
		VariableName: "POSTGRES_HOST",
	}

	POSTGRES_USER_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_USER"),
		VariableName: "POSTGRES_USER",
	}

	POSTGRES_PORT_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("POSTGRES_PORT"),
		VariableName: "POSTGRES_PORT",
	}

	PORT_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("PORT"),
		VariableName: "PORT",
	}

	DEBUG_OPTIONS := helpers.OptionsValidator{
		Value:        os.Getenv("DEBUG"),
		VariableName: "DEBUG",
	}

	return &env{
		POSTGRES_DB_NAME:  helpers.NewValidator(POSTGRES_DB_NAME_OPTIONS).Required().Text(),
		POSTGRES_DB_TABLE: helpers.NewValidator(POSTGRES_DB_TABLE_OPTIONS).Required().Text(),
		POSTGRES_DNS:      helpers.NewValidator(POSTGRES_DNS_OPTIONS).Required().Text(),
		POSTGRES_HOST:     helpers.NewValidator(POSTGRES_HOST_OPTIONS).Required().Text(),
		POSTGRES_USER:     helpers.NewValidator(POSTGRES_USER_OPTIONS).Required().Text(),
		POSTGRES_PORT:     helpers.NewValidator(POSTGRES_PORT_OPTIONS).Required().Text(),
		PORT:              helpers.NewValidator(PORT_OPTIONS).Required().Number(),
		DEBUG:             helpers.NewValidator(DEBUG_OPTIONS).Required().Boolean(),
	}
}
