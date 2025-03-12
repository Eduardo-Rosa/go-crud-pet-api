package handlers

import (
	"database/sql"
	"encoding/json"
	"go-crud-pet-api/models"
	"go-crud-pet-api/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repositories.CreatePet(db, pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetPets(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	pets, err := repositories.GetPets(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func GetPet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	pet, err := repositories.GetPet(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pet não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

func UpdatePet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet.ID = id
	if err := repositories.UpdatePet(db, pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeletePet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := repositories.DeletePet(db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
