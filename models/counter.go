package models

type Counter struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Counter int64 `gorm:"primaryKey" json:"counter"`
}
