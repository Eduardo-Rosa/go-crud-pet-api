package repositories

import (
	"database/sql"
	"go-crud-pet-api/models"
)

func CreatePet(db *sql.DB, pet models.Pet) error {
	_, err := db.Exec(`
        INSERT INTO pets (name, species, breed, age, birth_date, owner_name)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, pet.Name, pet.Species, pet.Breed, pet.Age, pet.BirthDate, pet.OwnerName)
	return err
}

func GetPets(db *sql.DB) ([]models.Pet, error) {
	rows, err := db.Query("SELECT * FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var pet models.Pet
		if err := rows.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Breed, &pet.Age, &pet.BirthDate, &pet.OwnerName); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func GetPet(db *sql.DB, id int) (models.Pet, error) {
	var pet models.Pet
	row := db.QueryRow("SELECT * FROM pets WHERE id = $1", id)
	err := row.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Breed, &pet.Age, &pet.BirthDate, &pet.OwnerName)
	return pet, err
}

func UpdatePet(db *sql.DB, pet models.Pet) error {
	_, err := db.Exec(`
        UPDATE pets
        SET name = $1, species = $2, breed = $3, age = $4, birth_date = $5, owner_name = $6
        WHERE id = $7
    `, pet.Name, pet.Species, pet.Breed, pet.Age, pet.BirthDate, pet.OwnerName, pet.ID)
	return err
}

func DeletePet(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM pets WHERE id = $1", id)
	return err
}
