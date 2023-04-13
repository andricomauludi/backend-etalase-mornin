package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-backend-etalase-mornin")) //membuka rute mysql (ini untuk local)
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{}) //melakukan migrate pada mysql

	DB = database
}
