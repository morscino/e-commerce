package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"e-commerce/config"
)

func connectPostgresDb(config config.ConfigType) *gorm.DB {
	dialect := postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=Africa/Lagos",
		config.PGHost,
		config.PGPort,
		config.PGUser,
		config.PGDatabase,
		config.PGPassword,
	))

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}

	return db
}
