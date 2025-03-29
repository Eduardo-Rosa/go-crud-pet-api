package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-crud-pet-api/handlers"

	"io/ioutil"
	_ "io/ioutil"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func runMigrations(db *sql.DB) {
	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		log.Fatalf("Erro ao ler o diretório de migrações: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile("./migrations/" + file.Name())
		if err != nil {
			log.Fatalf("Erro ao ler o arquivo de migração %s: %v", file.Name(), err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("Erro ao executar a migração %s: %v", file.Name(), err)
		}

		log.Printf("Migração %s executada com sucesso", file.Name())
	}
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	var err error
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	runMigrations(db)

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
