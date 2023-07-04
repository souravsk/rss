package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	postString := os.Getenv("PORT")
	if postString == "" {
		log.Fatal("PORT is not found in the .env")
	}
	fmt.Println("Port:", postString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	v1router.Get("/healthz", handlerRead)
	v1router.Get("/err",handlerError)

	router.Mount("/v1/",v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + postString,
	}

	log.Printf("server starting a port %v", postString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
