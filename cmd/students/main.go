package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Println("server running")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
