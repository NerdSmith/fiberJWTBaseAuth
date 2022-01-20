package config

import (
	"fiberJWTAuth/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DatabaseUri = "host=localhost user=postgres password=pgpwd dbname=postgres port=5432 sslmode=disable"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DatabaseUri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.BaseUser{})

	return nil
}
