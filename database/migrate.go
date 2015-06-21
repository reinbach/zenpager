package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/reinbach/zenpager/utils"
)

type Migration struct {
	App  string
	Name string
	File string
}

var (
	d      = utils.GetAbsDir()
	s      = map[string][]Migration{}
	c      = map[string][]Migration{}
	valid  = true
	preSql = regexp.MustCompile(`require [a-z]+ [0-9]{4}_[a-z]+.sql`)
)

func MigrationExists(app, name string) bool {
	// check that the name of migration does not
	// already exist in the list of found migrations
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
		App:  app,
		Name: filepath.Base(file),
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

func ExecSqlFiles(db *sql.DB, app string) error {
	log.Printf("%s:\n", app)
	for _, m := range s[app] {
		if err := ExecSqlFile(db, app, m); err != nil {
			return err
		}
	}
	return nil
}

func ExecSqlFile(db *sql.DB, app string, m Migration) error {
	if MigrationExists(app, m.Name) == false {
		if f, err := ioutil.ReadFile(m.File); err == nil {
			// check if there is a 'requirement'
			if err = ExecPreSqlFiles(db, f); err != nil {
				log.Println("error in pre exec...")
				return err
			}
			if err = ExecSql(db, f, m); err != nil {
				return err
			}
			c = MigrationAdd(c, app, m.Name)
		} else {
			log.Printf("Failed to get SQL from file %v: %v\n", m.File, err)
			return err
		}
	}
	return nil
}

func ExecPreSqlFiles(db *sql.DB, f []byte) error {
	// sql file for requirements
	// if any requirements, parse out app and name of migration
	// search migrations for match and then execute found match
	p := preSql.FindAll(f, -1)
	for _, n := range p {
		split := strings.Split(fmt.Sprintf("%s", n), " ")
		m, err := FindMigration(split[1], split[2])
		if err != nil {
			return err
		}
		if err = ExecSqlFile(db, split[1], m); err != nil {
			log.Println("Failed to run requirement: %v (%v)", split[1], m)
			return err
		}
	}
	return nil
}

func FindMigration(app, name string) (Migration, error) {
	for _, m := range s[app] {
		if m.Name == name {
			return m, nil
		}
	}
	return Migration{}, errors.New("Requirement File Not Found")
}

func ExecSql(db *sql.DB, f []byte, m Migration) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Not able to start transaction: ", err)
		os.Exit(1)
	}

	if _, err := tx.Exec(fmt.Sprintf("%s", f)); err != nil {
		log.Printf("Failed to migrate %v: %v\n", m.File, err)
		tx.Rollback()
		return err
	}
	log.Printf("- %s\n", m.Name)
	tx.Exec("INSERT INTO migrate (app, name, datetime) VALUES ($1, $2, NOW())",
		m.App, m.Name)

	tx.Commit()
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
			log.Printf("Issue scanning row: ", err)
			return err
		}
		c = MigrationAdd(c, app, name)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Issue iterating rows: ", err)
		return err
	}
	return nil
}

func Migrate() {
	filepath.Walk(d, FindSqlFiles)

	db := Connect()

	if err := GetRunSql(db); err != nil {
		log.Println("Failed to get previous migrations: ", err)
		os.Exit(1)
	}

	log.Println("Running migrations...\n")

	// 'database' migrations run first
	if err := ExecSqlFiles(db, "database"); err == nil {
		delete(s, "database")
		// Run the rest of the migrations
		for a, _ := range s {
			ExecSqlFiles(db, a)
		}
	}

	log.Println("\nFinished.\n")
}
