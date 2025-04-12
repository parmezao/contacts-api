package main

import (
	"contacts-api/database"
	"contacts-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/contacts", handlers.CreateContact).Methods("POST")
	r.HandleFunc("/contacts", handlers.GetContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", handlers.GetContact).Methods("GET")
	r.HandleFunc("/contacts/{id}", handlers.UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", handlers.DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
