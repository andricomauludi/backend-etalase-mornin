package models

type Klien struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaKlien  string `gorm:"type:varchar(100)" json:"nama_klien"`
	NomorHp    int64  `gorm:"type:varchar(30)" json:"nomor_hp"`
	EmailKlien int64  `gorm:"type:varchar(30)" json:"email_klien"`
}
