package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlshort/api"
	"urlshort/config"
	"urlshort/router"
	"urlshort/storage"
)

func main() {
	// linkStorage := storage.NewMemoryStorage()
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Info("Initializing storage...", "path", cfg.DBPath)
	linkStorage, err := storage.NewSqliteStorage(cfg.DBPath)
	if err != nil {
		logger.Error("failed to init storage", "error", err)
		os.Exit(1)
	}

	r := router.NewRouter()
	r.Handle("/", api.RedirectHandler(linkStorage))
	r.Handle("/api/v1/links", api.CreateLinkHandler(linkStorage))

	srv := &http.Server{
		Addr:         cfg.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("starting server", "port", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("could not start server", "error", err)
			os.Exit(1)
		}
	case sig := <-shutdownChan:
		logger.Info("shutting signal received", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("gracefully shutdown failed", "error", err)
			os.Exit(1)
		}
	}

	logger.Info("server stopped gracefully")
}
