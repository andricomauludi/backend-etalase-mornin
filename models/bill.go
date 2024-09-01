package models

import "time"

type Bill struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Id              int64  `gorm:"primaryKey" json:"id"`
	NamaBill        string `gorm:"type:varchar(100)" json:"nama_bill"`
	Paid            string `gorm:"type:varchar(1)" json:"paid"`
	Timestamp       time.Time
	JenisPembayaran string `gorm:"type:varchar(20)" json:"jenis_pembayaran"`
	Total           int64  `gorm:"type:integer(10)" json:"total"`
	CashIn          int64  `gorm:"type:integer(10)" json:"cash_in"`
	CashOut         int64  `gorm:"type:integer(10)" json:"cash_out"`
	IdKlien         int64  `gorm:"type:integer(13)" json:"id_klien"`
	NamaKlien       string `gorm:"type:varchar(20)" json:"nama_klien"`
	Tipe            string `gorm:"type:varchar(1)" json:"tipe"`
	//tipe 0 untuk ceu monny, tipe 1 untuk cvj
}
