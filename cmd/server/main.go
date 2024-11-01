package main

import (
	"log"

	"github.com/DJ-66/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8085")
	log.Fatal((srv.ListenAndServe()))
}

