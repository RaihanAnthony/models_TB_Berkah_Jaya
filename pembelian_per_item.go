package models

import (
	"gorm.io/gorm"
	"time"
)

type Pembelian_Per_Item struct {
	gorm.Model
	ID uint `gorm:"PrimaryKey" json:"ID"`
	PembelianID uint `json:"pembelianID" validate:"required"`
	BarangID uint `json:"barangID" validate:"required"`
	Barang Barang `gorm:"foreignKey:BarangID; references:ID" validate:"-"`
	// Pembelian Pembelian `gorm:"foreignKey:PembelianID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"UpdatedAt"`
	Jumlah_Barang float64 `gorm:"type:DECIMAL(10, 0);not null" json:"jumlah_barang" validate:"required"`
	Total_Harga float64 `gorm:"type:DECIMAL(10, 0);not null" json:"total_harga" validate:"required"`
	Total_Keuntungan float64 `gorm:"type:DECIMAL(10, 0);not null" json:"total_keuntungan" validate:"required"`
}