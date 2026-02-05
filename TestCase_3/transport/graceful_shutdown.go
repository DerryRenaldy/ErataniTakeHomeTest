package transport

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServer interface {
	GetState() ServerState
	SetState(state ServerState)
	GetGracePeriodSeconds() int64
	GetCleanupPeriodSeconds() int64
}

func SetupGracefulShutdown(httpServer HTTPServer) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go respondToSigterm(done, httpServer)
}

func respondToSigterm(done chan os.Signal, httpServer HTTPServer) {
	signal := <-done
	log.Printf("Received signal %v in process at %s. attempting graceful shutdown.", signal, time.Now())
	defer func() {
		log.Printf("Cleaning up completed. Shutting down now.")
		os.Exit(0)
	}()

	log.Printf("Entering grace period.")
	httpServer.SetState(ServerStateInGracePeriod)
	time.Sleep(time.Duration(httpServer.GetGracePeriodSeconds()) * time.Second)

	log.Printf("Entering cleanup period.")
	httpServer.SetState(ServerStateInCleanupPeriod)
	time.Sleep(time.Duration(httpServer.GetCleanupPeriodSeconds()) * time.Second)
}
