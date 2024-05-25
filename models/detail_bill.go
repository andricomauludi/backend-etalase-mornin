package models

type Detail_bill struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	// Id          int64  `gorm:"primaryKey" json:"id"`
	// ProductName string `gorm:"type:varchar(300)" json:"product_name"`
	// Price       int64  `gorm:"type:integer(10)" json:"price"`
	// Description string `gorm:"type:text" json:"description"`
	// MenuType    string `gorm:"type:varchar(300)" json:"menu_type"`
	// Photo       string `gorm:"type:varchar(300)" json:"photo"`
	Id         int64  `gorm:"primaryKey" json:"id"`
	IdBill     int64  `gorm:"type:integer(10)" json:"id_bill"`
	IdMenu     int64  `gorm:"type:integer(10)" json:"id_menu"`
	NamaMenu   string `gorm:"type:varchar(100)" json:"nama_menu"`
	Harga      int64  `gorm:"type:integer(10)" json:"harga"`
	Jumlah     int64  `gorm:"type:integer(10)" json:"jumlah"`
	TotalHarga int64  `gorm:"type:integer(10)" json:"total_harga"`
}
