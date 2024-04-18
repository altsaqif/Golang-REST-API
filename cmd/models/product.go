package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Nama   string  `gorm:"type:varchar(300);not null" json:"nama"`
	Stok   int32   `gorm:"type:int(5);not null" json:"stok"`
	Harga  float64 `gorm:"type:decimal(14,2);not null" json:"harga"`
	UserID uint    `gorm:"type:int(5);not null" json:"userid"`
}
