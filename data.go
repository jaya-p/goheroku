package goheroku

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// SetData will set field and its associated value
func SetData(field string, value string) (bool, error) {
	done := false

	db, errO := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errO != nil {
		log.Print(errO)
		return done, errO
	}

	// https://www.postgresqltutorial.com/postgresql-upsert/
	// https://pkg.go.dev/github.com/lib/pq?tab=doc
	_, errE := db.Exec("INSERT INTO data(field, value, update_ts) "+
		" VALUES($1, $2, CURRENT_TIMESTAMP) "+
		" ON CONFLICT (field) "+
		" DO UPDATE "+
		" SET field=$1, value=$2, update_ts=CURRENT_TIMESTAMP", field, value)
	if errE != nil {
		log.Print(errE)
		return done, errE
	}

	done = true

	return done, nil
}

// GetData will get value of a field
func GetData(field string) (string, error) {
	db, errO := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errO != nil {
		log.Print(errO)
		return "", errO
	}

	var value string
	var updateTs string
	row := db.QueryRow("SELECT value, update_ts FROM data WHERE field=$1", field)
	switch errR := row.Scan(&value, &updateTs); errR {
	case nil:
		return value + ", " + updateTs, nil
	default:
		log.Print(errR)
		return "", errR
	}
}
