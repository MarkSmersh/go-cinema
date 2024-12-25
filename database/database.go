package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MarkSmersh/go-cinema/scripts"
	_ "github.com/lib/pq"
)

var DB *sql.DB = initDatabase()

func initDatabase() *sql.DB {
	scripts.ParseDotEnv()

	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatalln(err.Error())
	}

	db.Ping()

	return db
}
