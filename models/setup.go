package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

<<<<<<< HEAD
var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-backend-etalase-mornin"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
=======
var DB *gorm.DB //menggunakan gorm db dalam koneksi db

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-backend-etalase-mornin")) //membuka rute mysql (ini untuk local)
	if err != nil {
		panic(err) //mengembalikan error apabila terdapat eror
	}

	database.AutoMigrate(&Product{}) //melakukan migrate pada mysql
>>>>>>> development

	DB = database
}
