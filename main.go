package main

import (
	"fmt"
	"github.com/MarkSmersh/go-cinema/api"
	"github.com/MarkSmersh/go-cinema/scripts"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	scripts.ParseDotEnv()

	fs := http.FileServer(http.Dir("./web/out/"))

	http.HandleFunc("/testik", api.Test)
	http.Handle("/", fs)

	fmt.Println("Server is started at port :8080")
	http.ListenAndServe(":8080", nil)
}
