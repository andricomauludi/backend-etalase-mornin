package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Harga       int64  `gorm:"type:integer(10)" json:"harga"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
