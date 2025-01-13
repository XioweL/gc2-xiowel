package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"time"
)

var DB *gorm.DB

// InitDB menginisialisasi koneksi database
func InitDB() {
	// Mengambil path untuk file .env
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../.env")

	// Memuat variabel environment dari .env
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load environment variables from %s: %v", envPath, err)
	}

	// Menyiapkan DSN (Data Source Name) untuk koneksi ke PostgreSQL
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai options='--client_encoding=UTF8' application_name=myapp",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Membuka koneksi database dengan preferensi debugging
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Nonaktifkan statement caching
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Mode debug untuk melihat query
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully.")

	// Mendapatkan instance SQL untuk pengaturan connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	// Mengatur connection pool
	sqlDB.SetMaxIdleConns(10)           // Jumlah koneksi idle maksimum
	sqlDB.SetMaxOpenConns(50)           // Jumlah koneksi terbuka maksimum
	sqlDB.SetConnMaxLifetime(time.Hour) // Waktu hidup koneksi

	DB = database
}

//// GetDB mengembalikan instance DB dengan session baru
//func GetDB() *gorm.DB {
//	return DB.Session(&gorm.Session{})
//}

// CloseDB menutup koneksi database
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance for closing: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
	log.Println("Database connection closed.")
}
