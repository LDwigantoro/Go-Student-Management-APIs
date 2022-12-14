package config

import (
	"log"
	"os"

	"github.com/LDwigantoro/go-clean-arch/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func DBConnect() *gorm.DB {

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal melakukan koneksi ke database" + dsn + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		entities.Mahasiswa{},
	)
	return db
}
