package database

import (
	. "ABIS/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSL")
	var err error
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Автоматическое создание таблиц
	err = db.AutoMigrate(&Book{}, &Author{})
	if err != nil {
		log.Fatalf("Не удалось создать таблицы: %v", err)
	}

	log.Println("Таблицы успешно созданы или уже существуют")
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
