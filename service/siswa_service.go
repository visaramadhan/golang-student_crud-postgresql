package service

import (
	"fmt"

	"github.com/go-embed-go-web/models"
	"github.com/go-embed-go-web/repository"
)

type SiswaService interface {
	RegisterNewSiswa(payload models.Siswa) error
	UpdateDataSiswa(siswa models.Siswa) error
	GetAllSiswa() ([]models.Siswa, error)
	DeleteSiswa(id int) error
}

type siswaService struct {
	repo repository.SiswaRepository
}

func (s *siswaService) DeleteSiswa(id int) error {
	return s.repo.Delete(id)
}

// update siswa implementasi SiswaService
func (s *siswaService) UpdateDataSiswa(siswa models.Siswa) error {
	return s.repo.Update(&siswa)
}

// getAll Siswa
func (s *siswaService) GetAllSiswa() ([]models.Siswa, error) {
	return s.repo.List()
}

// register new siswa implementasi SiswaService
func NewSiswaService(swarepo repository.SiswaRepository) SiswaService {
	return &siswaService{repo: swarepo}
}

func (s *siswaService) RegisterNewSiswa(payload models.Siswa) error {
	if payload.Nama == "" || payload.Class == "" {
		return fmt.Errorf("all fields in the payload are required")
	}
	err := s.repo.Create(&payload)
	if err != nil {
		return fmt.Errorf("failed to create siswa: %w", err)
	}
	return nil
}
