package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //menggunakan gorm db dalam koneksi db

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-backend-etalase-mornin?parseTime=true")) //membuka rute mysql (ini untuk local)
	if err != nil {
		panic(err) //mengembalikan error apabila terdapat eror
	}

	database.AutoMigrate(&Product{}, &User{}, &Transaction{}) //melakukan migrate pada mysql

	DB = database
	//test merge
}
