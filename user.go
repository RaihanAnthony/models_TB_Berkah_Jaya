package models

import( 
"gorm.io/gorm"
"time"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"UpdatedAt"` 
	UserName string `gorm:"type:varchar(100);unique;not null" json:"username" validate:"required,alphanum,uniqueUsername"`
	Email string	`gorm:"type:varchar(200);unique;not null" json:"email" validate:"required,email,uniqueEmail"`
	NoWhatshapp string `gorm:"type:varchar(20);unique;not null" json:"whatshapp" validate:"required,uniquePhone"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,alphanum,min=6"` 
	Poin float64   `gorm:"type:DECIMAL(10, 0)" json:"poin"`
	Pembelian []Pembelian 
	Hadiah  []Hadiah `gorm:"many2many:hadiah_users"`
}