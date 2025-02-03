package main

import (
	"context"
	"fmt"
	"go-lang-test-apis/internal/config"
	"go-lang-test-apis/internal/http/handlers/goapiuser"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Print("welcome\n")

	cnf := config.MustLoad()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /go/api/users", goapiuser.New())

	//setup server
	server := http.Server{
		Addr:    cnf.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cnf.HTTPserver.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error, but don't call log.Fatal which terminates the program
			slog.Error("Failed to start server", slog.String("error", err.Error()))
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
