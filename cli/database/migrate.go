package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"git.ironlabs.com/greg/zenpager/database"
	"git.ironlabs.com/greg/zenpager/utils"
)

type Migration struct {
	Name string
	File string
}

var (
	d     = utils.GetAbsDir()
	s     = map[string][]Migration{}
	c     = map[string][]Migration{}
	valid = true
)

func MigrationExists(app, name string) bool {
	if f, prs := c[app]; prs == true {
		for _, n := range f {
			if n.Name == name {
				return true
			}
		}
	}
	return false
}

func MigrationAdd(m map[string][]Migration, app, file string) map[string][]Migration {
	m[app] = append(m[app], Migration{
		Name: app,
		File: file,
	})
	return m
}

func FindSqlFiles(path string, info os.FileInfo, err error) error {
	base := filepath.Base(path)
	if base == "sql" {
		sql, err := filepath.Glob(filepath.Join(path, "*.sql"))
		if err != nil {
			return err
		}
		name := filepath.Base(path[:len(path)-len(base)])
		for _, f := range sql {
			s = MigrationAdd(s, name, f)
		}
	}
	return nil
}

func ExecSqlFiles(tx *sql.Tx, app string) error {
	fmt.Printf("%s:\n", app)
	for _, m := range s[app] {
		if MigrationExists(app, m.Name) == false {
			if f, err := ioutil.ReadFile(m.File); err == nil {
				if err = ExecSql(tx, f, m); err != nil {
					return err
				}
			} else {
				fmt.Printf("Failed to get SQL from file %v: %v\n", m.File, err)
				return err
			}
		}
	}
	return nil
}

func ExecSql(tx *sql.Tx, f []byte, m Migration) error {
	if _, err := tx.Exec(fmt.Sprintf("%s", f)); err != nil {
		fmt.Printf("Failed to migrate %v: %v\n", m.File, err)
		return err
	}
	fname := filepath.Base(m.File)
	fmt.Printf("- %s\n", fname)
	tx.Exec("INSERT INTO migrate (app, name, datetime) VALUES ($1, $2, NOW())",
		m.Name, fname)
	return nil
}

func GetRunSql(db *sql.DB) error {
	rows, err := db.Query("SELECT app, name FROM migrate")
	if err != nil {
		// table may not exist, so ignore this
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var app, name string
		if err := rows.Scan(&app, &name); err != nil {
			fmt.Printf("Issue scanning row: ", err)
			return err
		}
		c = MigrationAdd(c, app, name)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("Issue iterating rows: ", err)
		return err
	}
	return nil
}

func Migrate() {
	filepath.Walk(d, FindSqlFiles)

	db := database.Connect()

	if err := GetRunSql(db); err != nil {
		fmt.Println("Failed to get previous migrations: ", err)
		os.Exit(1)
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Not able to start transaction: ", err)
		os.Exit(1)
	}

	fmt.Println("Running migrations...\n")
	// Update `migrate` table first
	if err = ExecSqlFiles(tx, "database"); err != nil {
		valid = false
	} else {
		delete(s, "database")
		// Run the rest
		for a, _ := range s {
			ExecSqlFiles(tx, a)
		}
	}

	if valid == true {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	fmt.Println("\nFinished.\n")
}
