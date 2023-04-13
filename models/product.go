package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"varchar(300)" json:"nama_product"`
	Harga       int64  `gorm:"integer(10)" json:"harga"`
	Deskripsi   string `gorm:"varchar(30)" json:"deskripsi"`
}
