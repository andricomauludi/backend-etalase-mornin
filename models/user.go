package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(100)" json:"fullname"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"varchar(50)" json:"password"`
	Role     string `gorm:"type:varchar(5)" json:"role_id"`
}
