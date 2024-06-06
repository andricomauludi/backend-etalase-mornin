package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //menggunakan gorm db dalam koneksi db

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_PORT"))) //membuka rute mysql (ini untuk local)
	if err != nil {
		panic(err) //mengembalikan error apabila terdapat eror
	}

	database.AutoMigrate(&Product{}, &User{}, &Bill{}, &Detail_bill{}, &Klien{}, &Counter{}) //melakukan migrate pada mysql

	DB = database
	//test merge
}
