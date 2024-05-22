package models

type Product struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	// Id          int64  `gorm:"primaryKey" json:"id"`
	// ProductName string `gorm:"type:varchar(300)" json:"product_name"`
	// Price       int64  `gorm:"type:integer(10)" json:"price"`
	// Description string `gorm:"type:text" json:"description"`
	// MenuType    string `gorm:"type:varchar(300)" json:"menu_type"`
	// Photo       string `gorm:"type:varchar(300)" json:"photo"`
	Id            int64  `gorm:"primaryKey" json:"id"`
	JenisMenu     string `gorm:"type:varchar(100)" json:"jenis_menu"`
	NamaMenu      string `gorm:"type:text" json:"nama_menu"`
	DeskripsiMenu string `gorm:"type:text" json:"deskripsi_menu"`
	Harga         int64  `gorm:"type:integer(10)" json:"harga"`
	Photo         string `gorm:"type:varchar(300)" json:"photo"`
}
