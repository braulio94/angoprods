package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/category/{name}", Category).Methods("GET")
	router.HandleFunc("/api/product/{id}/", Product).Methods("GET")
	router.HandleFunc("", HomeRoute)
	router.HandleFunc("/", HomeRoute)
	router.NotFoundHandler = http.HandlerFunc(UnknownRoute)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
