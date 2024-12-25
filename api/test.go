package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarkSmersh/go-cinema/database"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(
		os.Getenv("DB_PORT"),
	)

	database.DB.Ping()

	rows, err := database.DB.Query("SELECT 1000")

	if err != nil {
		log.Fatalln(err)
	}

	println(rows)

	w.Write([]byte("It seems ok"))
}
