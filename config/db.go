package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=PostgreSQL user=mahdi password=aliali dbname=electronic port=5435 sslmode=disable",
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		println("Failed to connect database\"")
		os.Exit(1)
	}
}
