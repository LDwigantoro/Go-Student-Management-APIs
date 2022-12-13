package handler

import (
	"net/http"

	"github.com/LDwigantoro/go-clean-arch/entities"
	"github.com/LDwigantoro/go-clean-arch/usecases"
	"github.com/LDwigantoro/go-clean-arch/utils"
	"github.com/gin-gonic/gin"
)

type MahasiswaHandler struct {
	mahasiswaUsecase usecases.IMahasiswaUsecase
}

func CreateMahasiswaHandler(r *gin.Engine, mahasiswaUsecase usecases.IMahasiswaUsecase) {
	mahasiswaHandler := MahasiswaHandler{mahasiswaUsecase}

	r.POST("/mahasiswa", mahasiswaHandler.AddMahasiswa)
}

func (h *MahasiswaHandler) AddMahasiswa(c *gin.Context) {
	var mahasiswa = entities.Mahasiswa{}

	err := c.Bind(&mahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Server Error")
		return
	}

	// jika nama mahasiswa kosong
	if mahasiswa.FirstName == "" || mahasiswa.LastName == "" {
		utils.HandleError(c, http.StatusBadRequest, "Nama tidak boleh kosong")
		return
	}

	newMahasiswa, err := h.mahasiswaUsecase.Create(&mahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	utils.HandleSuccess(c, newMahasiswa)
}
