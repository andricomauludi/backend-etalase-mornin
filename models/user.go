package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(100)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(50)" json:"username"`
	Password    string `gorm:"varchar(50)" json:"password"`
}
