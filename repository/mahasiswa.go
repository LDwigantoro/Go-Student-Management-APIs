package repository

import (
	"fmt"

	"github.com/LDwigantoro/go-clean-arch/entities"
	"gorm.io/gorm"
)

type IMahasiswaRepository interface {
	Create(mahasiwa *entities.Mahasiswa) (*entities.Mahasiswa, error)
	ReadAll() (*[]entities.Mahasiswa, error)
}

type MahasiswaRepository struct {
	DB *gorm.DB
}

func NewMahasiswaRepository(DB *gorm.DB) IMahasiswaRepository {
	return &MahasiswaRepository{DB}
}

func (m *MahasiswaRepository) Create(mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error) {
	err := m.DB.Save(&mahasiswa).Error
	if err != nil {
		fmt.Printf("[MahasiswaRepository.Create] gagal melakukan eksekusi query %v \n", err)
		return nil, fmt.Errorf("gagal memasukan data")
	}
	return mahasiswa, nil

}

func (m *MahasiswaRepository) ReadAll() (*[]entities.Mahasiswa, error) {
	var mahasiswa []entities.Mahasiswa
	err := m.DB.Find(&mahasiswa).Error

	if err != nil {
		fmt.Printf("[MahasiswaRepository.ReadAll] gagal melakukan eksekusi query %v \n", err)
		return nil, fmt.Errorf("gagal mengeluarkan semua data mahasiswa")
	}

	return &mahasiswa, nil
}
