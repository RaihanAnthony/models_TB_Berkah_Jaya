package models

import (
	"gorm.io/gorm"
	"time"
)

type Hadiah struct {
	gorm.Model
	ID  uint `gorm:"id" json:"ID"` 
	CreatedAt time.Time `gorm:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"UpdatedAt"`
	Nama_Barang string `gorm:"type:varchar(200);not null" json:"nama_barang" validate:"required"`
	Harga_Hadiah float64 `gorm:"type:DECIMAL(10,0);not null" json:"harga_hadiah" validate:"required, number`
	Poin float64 `gorm:"type:DECIMAL(10,0);not null" json:"poin" validate:"required,number"`
	Image string `gorm:"type:varchar(200); not null" json:"image" validate:"required"`
	Deskripsi string `gorm:"type:varchar(500)" json:"desc" validate:"max=500"`
	User []User `gorm:"many2many:hadiah_users"`
}

type GetHadiah struct{
	Nama_Barang string `gorm:type:varchar(100);not null" json:"nama_barang`
	Poin int64
	Image string
	Desc string
}