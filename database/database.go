package database

import (
	"database/sql"
	"log"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	_ "github.com/lib/pq"
	"github.com/zenazn/goji/web"
)

const (
	DB_KEY string = "database"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", GetDatasource())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DB_KEY, db)
}

func FromContext(c web.C) *sql.DB {
	ctx := webctx.FromC(c)
	return ctx.Value(DB_KEY).(*sql.DB)
}
