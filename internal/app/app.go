package app

import (
	"context"
	"github.com/forabelo/test-task-perl/config"
	"github.com/forabelo/test-task-perl/internal/app/deps"
	"github.com/forabelo/test-task-perl/internal/infrastructure/http_server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	// initialize dependencies
	dependencies, err := deps.Init(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// initialize http server
	httpServer, err := http_server.NewHTTPServer(dependencies)
	if err != nil {
		log.Fatalf("Failed to initialize http server: %v", err)
	}

	// run http server
	err = httpServer.Run(":" + string(cfg.ServerPort))
	if err != nil {
		log.Fatalf("Failed to initialize http server: %v", err)
	}

	// graceful shutdown code
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// add handling panics
	go func() {
		<-shutdown

		os.Exit(0)
	}()

	<-shutdownCtx.Done()
}
