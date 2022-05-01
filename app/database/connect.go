package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"travel-db.com", 5432, user, dbname, password,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      newLogger,
		},
	)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		Account{},
		Comment{},
		Marker{},
		Post{},
		Request{},
		Friend{},
		Mute{},
		Like{},
		Session{},
	)

	return db
}
