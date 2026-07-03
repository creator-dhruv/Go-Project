package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/creator-dhruv/Go-Project/internal/config"
	"github.com/creator-dhruv/Go-Project/internal/http/routes"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Database setup

	// Router setup
	router := http.NewServeMux()
	routes.UserRouter(router)

	// Server setup
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("Server started", slog.String("Address", cfg.Address))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting Down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to Shutdown", slog.String("error : ", err.Error()))
	}

	slog.Info("Server ShutDown successfully")

}
