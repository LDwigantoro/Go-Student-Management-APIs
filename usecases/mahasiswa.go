package usecases

import (
	"github.com/LDwigantoro/go-clean-arch/entities"
	"github.com/LDwigantoro/go-clean-arch/repository"
)

type IMahasiswaUsecase interface {
	Create(mahasiwa *entities.Mahasiswa) (*entities.Mahasiswa, error)
	ReadAll() (*[]entities.Mahasiswa, error)
	Read(id int) (*entities.Mahasiswa, error)
	Update(id int, mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error)
}

type MahasiswaUsecase struct {
	mahasiswaRepo repository.IMahasiswaRepository
}

func NewMahasiswaUsecase(mahasiswaRepo repository.IMahasiswaRepository) IMahasiswaUsecase {
	return &MahasiswaUsecase{mahasiswaRepo}
}

func (u *MahasiswaUsecase) Create(mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error) {
	return u.mahasiswaRepo.Create(mahasiswa)
}

func (u *MahasiswaUsecase) ReadAll() (*[]entities.Mahasiswa, error) {
	return u.mahasiswaRepo.ReadAll()
}

func (u *MahasiswaUsecase) Read(id int) (*entities.Mahasiswa, error) {
	return u.mahasiswaRepo.Read(id)
}

func (u *MahasiswaUsecase) Update(id int, mahasiswa *entities.Mahasiswa) (*entities.Mahasiswa, error) {
	return u.mahasiswaRepo.Update(id, mahasiswa)
}
