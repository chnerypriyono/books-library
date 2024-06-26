package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
	"context"
	"firebase.google.com/go/v4/auth"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *slog.Logger
	authClient *auth.Client
}

func main() {

	var cfg config

	// Try to read environment variable for port (given by railway). Otherwise use default
	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 4000
	}

	// Set the port to run the API on
	cfg.port = intPort

	// create the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	authClient := initFirebaseAuth(context.Background(), logger)

	// create the application
	app := &application{
		config: cfg,
		logger: logger,
		authClient: authClient,
	}

	// create the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("server started", "addr", srv.Addr)

	// Start the server
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
