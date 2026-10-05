package config

import (
	"event-app/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		log.Fatal("Environment Variabel Belum diisi")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal Terkoneksi di Database", err)
	}

	err = database.AutoMigrate(&models.Event{}, models.User{})
	if err != nil {
		log.Fatal("Gagal Melakukan migration database", err)
	}

	DB = database
	log.Println("Berhasil Terkoneksi ke database")
}
