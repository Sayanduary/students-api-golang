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

	"github.com/Sayanduary/students-api-golang.git/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// db setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}
	slog.Info("sever started", slog.String("Address:", cfg.HttpServer.Addr))

	// need to implement graceful shutdown
	// for ongoing process (first process the request then stop)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shuting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfuly")
}
