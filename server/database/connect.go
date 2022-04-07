package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"travel-db.com", 5432, user, dbname, password,
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		Comments{},
		Markers{},
		Posts{},
		Requests{},
		Users{},
		Likes{},
		Friends{},
		Mutes{},
	)

	return db
}
