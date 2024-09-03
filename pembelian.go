package models

import (
	"gorm.io/gorm"
	"time"
)

type Pembelian struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	UserID uint `json:"userid" validate:"required"`
	User User `gorm:"foreignKey:UserID;references:ID" validate:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"UpdatedAt"`
	Tanggal_Pembelian  string `gorm:"type:varchar(200)" json:"tanggal_pembelian" gorm:"required,alpha"`
	Total_Harga float64 `gorm:"type:DECIMAL(10, 0)" json:"total_harga" validate:"required"`
	Total_Keuntungan float64 `gorm:"type:DECIMAL(10, 0)" json:"total_keuntungan" validate:"required"`
	Image string `gorm:"type:varchar(300); not null" json:"image" validate:"required"`
	KeteranganNotaCancelID uint `json:"kerangan_nota_cancel_id"`
	Status string `gorm:"type:enum('cancel', 'success');default:'cancel'"`
	KeteranganNotaCancel KeteranganNotaCancel `gorm:"foreignKey:KeteranganNotaCancelID;references:ID"`
}

type KeteranganNotaCancel struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"UpdatedAt"`
}

type PembelianResponse struct {
	ID               uint      `json:"ID"`
	CreatedAt        time.Time `json:"CreatedAt"`
	UpdatedAt        time.Time `json:"UpdatedAt"`
	UserID           uint      `json:"userid"`
	UserName		 string    `json:"username"`
	Email			 string    `json:"email"`
	Tanggal_Pembelian  string  `json:"tanggal_pembelian"`
	Total_Harga      float64   `json:"total_harga"`
	Total_Keuntungan float64   `json:"total_keuntungan"`
	Image            string    `json:"image"`
	Keterangan		string	   `json:"keterangan"`
}

