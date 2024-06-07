package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //menggunakan gorm db dalam koneksi db

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_PORT")
	if dsn == "" {
		log.Fatal("DATABASE_PORT environment variable is not set")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = database.AutoMigrate(&Product{}, &User{}, &Bill{}, &Detail_bill{}, &Klien{}, &Counter{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = database
	log.Println("Database connection established and migration completed.")
}
