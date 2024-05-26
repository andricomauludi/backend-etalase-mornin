package models

type Klien struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaKlien  string `gorm:"type:varchar(100)" json:"nama_klien"`
	NomorHp    string `gorm:"type:varchar(30)" json:"nomor_hp"`
	EmailKlien string `gorm:"type:varchar(30)" json:"email_klien"`
}
