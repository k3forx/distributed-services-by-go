package main

import (
	"log"

	"github.com/k3forx/distributed-services-by-go/LetsGo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
