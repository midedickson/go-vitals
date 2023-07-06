package main

import (
	"log"

	"github.com/Double-DOS/go-vitals/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
