package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"test/model"
)

var DB *gorm.DB

func ConnectDB() {

    // Setting untuk masuk ke database dimasukkan dalam variable source
	source := "host=localhost user=your_username password=your_password dbname=your_dbname port=5432"

	// Membuat koneksi ke database menggunakan GORM
	database, err := gorm.Open(postgres.Open(source), &gorm.Config{})

	// Error handling jika koneksi ke database gagal
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	// Log jika koneksi ke database berhasil
	log.Println("Database connected")
	database.AutoMigrate(&model.Post{})

	// Menyimpan koneksi database ke variable DB
	DB = database
}
