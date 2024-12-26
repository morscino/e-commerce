package db

import (
	"gorm.io/gorm"

	"e-commerce/config"
)

type Database struct {
	PostgresDb *gorm.DB
}

func ConnectDB(config config.ConfigType) Database {
	db := Database{
		PostgresDb: connectPostgresDb(config),
	}
	return db
}
