package main

import (
	_ "github.com/MarkSmersh/go-cinema/init"

	"fmt"
	"github.com/MarkSmersh/go-cinema/api"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fs := http.FileServer(http.Dir("./web/out/"))

	http.Handle("/", fs)
	http.HandleFunc("/testik", api.Test)

	fmt.Println("Server is started at port :8080")
	http.ListenAndServe(":8080", nil)
}
