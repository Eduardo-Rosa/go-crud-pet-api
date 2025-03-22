package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"), os.Getenv("DB_SSLMODE"))

	var err error
	for i := 0; i < 10; i++ { // Tenta conectar 10 vezes
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Println("Tentando conectar ao banco de dados... tentativa:", i+1)
		time.Sleep(3 * time.Second) // Espera antes de tentar novamente
	}
	return nil, err
}

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
