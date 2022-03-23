package main

import (
	"cmd/main/main.go/internal/user"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userHandler := user.NewHandler()
	userHandler.Register(router)

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: router,
	}

	fmt.Println("Server has been started")
	log.Fatal(server.Serve(listener))
}
