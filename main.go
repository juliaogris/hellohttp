package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	lis := "localhost:8080"
	if len(os.Args) > 1 {
		lis = os.Args[1]
	}
	fmt.Println("starting server on", lis)
	log.Fatal(http.ListenAndServe(lis, nil))
}
