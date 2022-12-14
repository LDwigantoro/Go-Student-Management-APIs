package repository

import (
	"fmt"

	"github.com/LDwigantoro/go-clean-arch/entities"
	"github.com/jinzhu/gorm"
)

type IMahasiswaRepository interface {
	Create(mahasiwa *entities.Mahasiswa) (*entities.Mahasiswa, error)
	ReadAll() (*[]entities.Mahasiswa, error)
	Read(id int) (*entities.Mahasiswa, error)
	Update(id int, mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error)
	Delete(id int) error
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

func (m *MahasiswaRepository) Read(id int) (*entities.Mahasiswa, error) {
	var mahasiswa = entities.Mahasiswa{}

	err := m.DB.Table("mahasiswa").Where("id = ?", id).First(&mahasiswa).Error

	if err != nil {
		fmt.Printf("[MahasiswaRepository.Read] gagal melakukan eksekusi query %v \n", err)
		return nil, fmt.Errorf("data mahasiswa tidak ada")
	}
	return &mahasiswa, nil
}

func (m *MahasiswaRepository) Update(id int, mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error) {
	var updateMahasiswa = entities.Mahasiswa{}

	err := m.DB.Table("mahasiswa").Where("id = ?", id).First(&updateMahasiswa).Update(&mahasiswa).Error

	if err != nil {
		fmt.Printf("[MahasiswaRepository.Update] gagal melakukan eksekusi query %v \n", err)
		return nil, fmt.Errorf("data mahasiswa tidak bisa diupdate")
	}

	return &updateMahasiswa, nil
}

func (m *MahasiswaRepository) Delete(id int) error {
	var deleteMahasiswa = entities.Mahasiswa{}

	err := m.DB.Table("mahasiswa").Where("id = ?", id).First(&deleteMahasiswa).Delete(&deleteMahasiswa).Error

	if err != nil {
		fmt.Printf("[MahasiswaRespository.Delete] gagal melakukan eksekusi query %v \n", err)
		return fmt.Errorf("data mahasiswa tidak bisa ditemukan")
	}
	return nil
}
