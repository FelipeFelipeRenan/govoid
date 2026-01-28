package main

import (
	"log/slog"
	"os"

	"github.com/FelipeFelipeRenan/govoid/internal/engine"
	"github.com/FelipeFelipeRenan/govoid/internal/transport"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	storage := engine.NewStringStore()

	server := transport.New(":9000", storage, logger)
	
	logger.Info("GoVoid activating")
	if err := server.Start(); err != nil{
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}