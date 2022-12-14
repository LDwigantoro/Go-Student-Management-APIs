package handler

import (
	"net/http"
	"strconv"

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
	r.GET("/mahasiswa", mahasiswaHandler.GetAllMahasiswa)
	r.GET("/mahasiswa/:id", mahasiswaHandler.GetMahasiswa)
	r.PUT("/mahasiswa/:id", mahasiswaHandler.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", mahasiswaHandler.HapusMahasiswa)
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

func (h *MahasiswaHandler) GetAllMahasiswa(c *gin.Context) {
	mahasiswa, err := h.mahasiswaUsecase.ReadAll()

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Server Error")
		return
	}

	// jika tidak ada data mahasiwa di database
	if len(*mahasiswa) == 0 {
		utils.HandleError(c, http.StatusBadGateway, "Tidak ada data mahasiswa di database")
	}

	utils.HandleSuccess(c, mahasiswa)
}

func (h *MahasiswaHandler) GetMahasiswa(c *gin.Context) {

	idMahasiswa := c.Param("id")
	id, err := strconv.Atoi(idMahasiswa)

	if err != nil {

		utils.HandleError(c, http.StatusBadGateway, "ID Tidak ditemukan")
		return
	}

	mahasiswa, err := h.mahasiswaUsecase.Read(id)

	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
	}

	utils.HandleSuccess(c, mahasiswa)
}

func (h *MahasiswaHandler) UpdateMahasiswa(c *gin.Context) {
	idMahasiswa := c.Param("id")

	id, err := strconv.Atoi(idMahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusBadGateway, "ID Tidak ditemukan")
		return
	}

	_, err = h.mahasiswaUsecase.Read(id)

	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	var tempMahasiswa = entities.Mahasiswa{}
	err = c.Bind(&tempMahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Server Error")
		return
	}

	if tempMahasiswa.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "ID Salah")
		return
	}

	if tempMahasiswa.FirstName == "" || tempMahasiswa.LastName == "" {
		utils.HandleError(c, http.StatusBadRequest, "Nama tidak boleh kosong")
		return
	}

	mahasiswa, err := h.mahasiswaUsecase.Update(id, &tempMahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, mahasiswa)

}

func (h *MahasiswaHandler) HapusMahasiswa(c *gin.Context) {
	idMahasiswa := c.Param("id")

	id, err := strconv.Atoi(idMahasiswa)

	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "ID salah")
		return
	}

	err = h.mahasiswaUsecase.Delete(id)

	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccess(c, "Berhasil menghapus data mahasiswa")

}
