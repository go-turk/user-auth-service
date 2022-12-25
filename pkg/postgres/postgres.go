package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(setDSN()), &gorm.Config{})
	return db, err
}

func setDSN() string {
	host := "localhost"
	if data := os.Getenv("POSTGRES_HOST"); data != "" {
		host = data
	}
	port := "5432"
	if data := os.Getenv("POSTGRES_PORT"); data != "" {
		port = data
	}
	user := "postgres"
	if data := os.Getenv("POSTGRES_USER"); data != "" {
		user = data
	}
	password := "postgres"
	if data := os.Getenv("POSTGRES_PASSWORD"); data != "" {
		password = data
	}
	dbname := "postgres"
	if data := os.Getenv("POSTGRES_DB"); data != "" {
		dbname = data
	}
	return "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable TimeZone=Europe/Istanbul"
}
