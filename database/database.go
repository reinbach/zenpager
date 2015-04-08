package database

import (
	"database/sql"
	"fmt"
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

func InitDB() {
	Migrate()
}

func DropTables() {
	db := Connect()

	rows, err := db.Query(
		"SELECT table_name FROM information_schema.tables WHERE table_schema = $1",
		"public",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err = rows.Scan(&table); err != nil {
			log.Fatal(err)
		}
		stmt := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
