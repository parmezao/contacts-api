package handlers

import (
	"contacts-api/database"
	"contacts-api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	json.NewDecoder(r.Body).Decode(&contact)

	database.DB.Create(&contact)
	json.NewEncoder(w).Encode(contact)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact
	database.DB.Find(&contacts)

	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var contact models.Contact

	id, _ := strconv.Atoi(vars["id"])
	database.DB.First(&contact, id)
	json.NewEncoder(w).Encode(contact)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var contact models.Contact
	database.DB.First(&contact, id)

	json.NewDecoder(r.Body).Decode(&contact)
	database.DB.Save(&contact)
	json.NewEncoder(w).Encode(contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var contact models.Contact
	database.DB.Delete(&contact, id)
	w.WriteHeader(http.StatusNoContent)
}
