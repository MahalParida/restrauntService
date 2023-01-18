package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(username, password, dbName string) *sqlx.DB {

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(db:3306)/%s?parseTime=true", username, password, dbName))

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
