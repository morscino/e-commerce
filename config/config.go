package config

import (
	"e-commerce/helpers"
	"errors"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

type ConfigType struct {
	AppEnv          string
	Port            string
	JwtSecret       string `validate:"required"`
	JwtSecretExpiry string `validate:"required"`
	AppHost         string
	EnableSwagger   string
	PGHost          string `validate:"required"`
	PGPort          string `validate:"required"`
	PGUser          string `validate:"required"`
	PGPassword      string `validate:"required"`
	PGDatabase      string `validate:"required"`
}

func GetConfig() *ConfigType {
	if os.Getenv("APP_ENV") != "prod" && os.Getenv("APP_ENV") != "stg" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal().Err(err).Msgf("env file error: %s", err.Error())
		}
	}

	ConfigVariables := ConfigType{
		AppEnv:          helpers.Getenv("APP_ENV", "local"),
		Port:            helpers.Getenv("PORT", "7000"),
		JwtSecret:       helpers.Getenv("JWT_SECRET", "HYUOH65432FGJUYTRE5GTYUO"),
		JwtSecretExpiry: helpers.Getenv("JWT_SECRET_EXPIRY", "3h"),
		EnableSwagger:   helpers.Getenv("ENABLE_SWAGGER", "true"),
		AppHost:         helpers.Getenv("APP_HOST", "0.0.0.0"),
		PGHost:          os.Getenv("PG_HOST"),
		PGPort:          os.Getenv("PG_PORT"),
		PGUser:          os.Getenv("PG_USER"),
		PGPassword:      os.Getenv("PG_PASSWORD"),
		PGDatabase:      os.Getenv("PG_DATABASE"),
	}

	errs := helpers.ValidateInput(ConfigVariables)

	if len(errs) > 0 {
		for _, err := range errs {
			log.Fatal().Err(errors.New(err)).Msgf("env validation error: %s", err)
		}

	}
	return &ConfigVariables
}
