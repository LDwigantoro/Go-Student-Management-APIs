package main

import (
	"log"

	"github.com/LDwigantoro/go-clean-arch/config"
	mahasiswaHandler "github.com/LDwigantoro/go-clean-arch/handler"
	mahasiswaRepo "github.com/LDwigantoro/go-clean-arch/repository"
	mahasiswaUsecase "github.com/LDwigantoro/go-clean-arch/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "9090"
	db := config.DBConnect()

	router := gin.Default()

	mahasiswaRepo := mahasiswaRepo.NewMahasiswaRepository(db)
	mahasiswaUsecase := mahasiswaUsecase.NewMahasiswaUsecase(mahasiswaRepo)
	mahasiswaHandler.CreateMahasiswaHandler(router, mahasiswaUsecase)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
