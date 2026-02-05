package main

import (
	"log"

	"eratani_assesment_test/TestCase_3/transport"
)

//go:generate go run github.com/google/wire/cmd/wire

func main() {
	log.Println("Starting application...")

	// Wire initialization
	httpSvc := InitializeHTTPService()

	// Global graceful shutdown for http server
	transport.SetupGracefulShutdown(httpSvc)

	// Run server
	httpSvc.SetupAndServe()
}
