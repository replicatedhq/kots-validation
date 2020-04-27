package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MustGetDBSession() *sql.DB {
	if DB != nil {
		return DB
	}

	log.Printf("Connecting to db at %s", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("error connecting to mysql: %v\n", err)
		panic(err)
	}

	DB = db

	return db
}
