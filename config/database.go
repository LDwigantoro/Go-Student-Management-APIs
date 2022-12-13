package config

import (
	"log"
	"os"

	"github.com/LDwigantoro/go-clean-arch/entities"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := os.Getenv("DB_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal melakukan koneksi ke database" + dsn + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		entities.Mahasiswa{},
	)
	return db
}
