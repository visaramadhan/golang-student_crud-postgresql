package repository

import (
	"database/sql"
	// "fmt"

	"github.com/go-embed-go-web/models"
)

// Repository interface for managing packages
type SiswaRepository interface {
	Create(*models.Siswa) error
	Update(*models.Siswa) error
	Delete(id int) error
	// GetByID(id int) (*models.Siswa, error)
	List() ([]models.Siswa, error)
}

// SiswaRepository struct
type siswaRepository struct {
	db *sql.DB
}

// NewSiswaRepository creates a new SiswaRepository instance

func NewSiswaRepository(db *sql.DB) SiswaRepository {
	return &siswaRepository{db: db}
}

// Create creates a new Siswa

func (r *siswaRepository) Create(siswa *models.Siswa) error {
	query := "INSERT INTO siswa (nama, kelas) VALUES ($1,$2)"
	_, err := r.db.Exec(query, siswa.Nama, siswa.Class)
	return err
}

// Update updates an existing Siswa

func (r *siswaRepository) Update(siswa *models.Siswa) error {
	query := "UPDATE siswa SET nama=$1, kelas=$2 WHERE siswa_id=$3"
	_, err := r.db.Exec(query, siswa.Nama, siswa.Class, siswa.ID)
	return err
}

// Delete deletes a Siswa by ID

func (r *siswaRepository) Delete(id int) error {
	query := "DELETE FROM siswa WHERE siswa_id=$1"
	_, err := r.db.Exec(query, id)
	return err
}

// // GetByID retrieves a Siswa by ID

// func (r *siswaRepository) GetByID(id int) (*models.Siswa, error) {
// 	query := "SELECT siswa_id, nama, kelas FROM siswa WHERE id=?"
// 	row := r.db.QueryRow(query, id)
// 	siswa := &models.Siswa{}
// 	err := row.Scan(&siswa.ID, &siswa.Nama, &siswa.Class)
// 	if err == sql.ErrNoRows {
// 		return nil, fmt.Errorf("siswa not found with ID %d", id)
// 	}
// 	return siswa, err
// }

// List retrieves all Siswa

func (r *siswaRepository) List() ([]models.Siswa, error) {
	query := "SELECT siswa_id, nama, kelas FROM siswa"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var siswas []models.Siswa
	for rows.Next() {
		siswa := &models.Siswa{}
		err := rows.Scan(&siswa.ID, &siswa.Nama, &siswa.Class)
		if err != nil {
			return nil, err
		}
		siswas = append(siswas, *siswa)
	}
	return siswas, err
}
