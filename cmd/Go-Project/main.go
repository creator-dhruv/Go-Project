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
	"github.com/creator-dhruv/Go-Project/internal/storage/sqlite"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal("database is not connected : ", err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// Router setup
	router := http.NewServeMux()
	routes.UserRouter(router, storage)

	// Server setup
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("server started", slog.String("address", cfg.Address))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("error : ", err.Error()))
	}

	slog.Info("server shutDown successfully")

}
