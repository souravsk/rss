package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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
