package entities

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sanyuanya/doctor/config"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("postgres", config.DATABASE_URL)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func Transaction() (*sql.Tx, error) {
	return db.Begin()
}
