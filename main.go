package main

import (
	"log"
)

func main() {
	srv := NewCatFactService("https://catfact.ninja/fact")
	srv = NewLoggingService(srv)

	apiServer := NewApiServer(srv)
	log.Fatal(apiServer.Start(":3000"))
}
