package models

import (
	"gorm.io/gorm"
	"time"
)

type Barang struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at`
	Nama_Barang  string  `gorm:"Varchar(300); not null; unique" json:"nama_barang" validate:"required"`
	Harga_Barang  float64  `gorm:"type:DECIMAL(10,0);not null" json:"harga_barang" validate:"required,numeric"`
	Harga_Beli float64 `gorm:"type:DECIMAL(10,0);not null" json:"harga_beli" validate:"required,numeric"`
	Image string `gorm:"varchar(200)" json:"image"`
	Kode string  `gorm:"type:varchar(200); unique" json:"kode" validate:"number"`
}