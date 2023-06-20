package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NamaLengkap string `gorm:"varchar(100)" json:"nama_lengkap"`
	Username    string `gorm:"unique" json:"username"`
	Password    string `gorm:"varchar(50)" json:"password"`
	Role        string `gorm:"varchar(50)" json:"role_id"`
}
