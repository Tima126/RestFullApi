package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// DSN берём из env
	dsn := os.Getenv("DB_DSN")
	fmt.Println("Using DSN:", dsn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go server! DSN: %s", dsn)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
