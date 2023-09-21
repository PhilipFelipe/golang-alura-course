package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/PhilipFelipe/golang-alura-course/database"
	"github.com/PhilipFelipe/golang-alura-course/models"
	"github.com/gorilla/mux"
)

func ListPersonalities(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetrievePersonality(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(&newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(&personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(&personality)
}
