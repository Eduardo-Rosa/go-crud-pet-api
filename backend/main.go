package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"go-crud-pet-api/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Conectar ao banco de dados
	connStr := "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" host=" + os.Getenv("DB_HOST") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configurar o roteador
	r := mux.NewRouter()

	// Rotas
	r.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePet(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPets(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/pets/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPet(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/pets/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePet(w, r, db)
	}).Methods("PUT")
	r.HandleFunc("/pets/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePet(w, r, db)
	}).Methods("DELETE")

	// Iniciar o servidor
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
