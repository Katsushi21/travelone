package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/Katsushi21/travelone/ent"
	"github.com/Katsushi21/travelone/ent/migrate"
	_ "github.com/lib/pq"
)

func Connect() *ent.Client {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"travel-db.com", 5432, user, dbname, password,
	)

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	return client
}
