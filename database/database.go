package database

import (
	"database/sql"
	"log"

	"code.google.com/p/go.net/context"
	_ "github.com/lib/pq"
)

const (
	DB_KEY string = "database"
)

func Connect(datasource string) *sql.DB {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DB_KEY, db)
}

func FromContext(ctx context.Context) *sql.DB {
	db, ok := ctx.Value(DB_KEY).(sql.DB)
	if !ok {
		log.Fatalf("Expected Database connecting in context, got: %v", db)
	}
	return &db
}
