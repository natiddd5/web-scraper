package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT not found in enviorment")
	}
	fmt.Println(portString)

	router := chi.NewRouter()
	//Allow requests to HTTP / HTTPS, allowed methods are GET POST PUT DELETE OPTIONS.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handleError)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Println("Server running on port " + portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
