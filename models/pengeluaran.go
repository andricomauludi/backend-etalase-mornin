package models

import "time"

type Pengeluaran struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Id               int64  `gorm:"primaryKey" json:"id"`
	NamaPengeluaran  string `gorm:"type:varchar(200)" json:"nama_pengeluaran"`
	JenisPengeluaran string `gorm:"type:varchar(50)" json:"jenis_pengeluaran"`
	WaktuPengeluaran time.Time
	HargaPengeluaran int64  `gorm:"type:integer(10)" json:"harga_pengeluaran"`
	JumlahBarang     int64  `gorm:"type:integer(10)" json:"jumlah_barang"`
	Satuan           string `gorm:"type:varchar(50)" json:"satuan"`
	TotalPengeluaran int64  `gorm:"type:integer(10)" json:"total_pengeluaran"`
}
