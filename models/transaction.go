package models

import "time"

type Transaction struct { //CATATAN PENTING, variable row harus dimulai dengan HURUF KAPITAL atau ngga eror
	Id                int64     `gorm:"primaryKey" json:"id"`
	NumberTransaction string    `gorm:"type:integer(3)" json:"number_transaction"`
	Timestamp         time.Time `json:"timestamp"`
	Items             int64     `gorm:"type:text" json:"items"`
	ItemCount         int64     `gorm:"type:integer(3)" json:"item_count"`
	TotalPrice        int64     `gorm:"type:integer(3)" json:"total_price"`
	PaymentWith       string    `gorm:"type:varchar(30)" json:"payment_with"`
	Cashier           string    `gorm:"type:varchar(50)" json:"cashier"`
}
