package domain

import (
	"database/sql"
)

type Product struct {
	ID          int          `gorm:"column:id;primaryKey"`
	CreatedAt   sql.NullTime `gorm:"column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"column:updated_at"`
	Category    string       `gorm:"column:category"`
	Price       int64        `gorm:"column:price"`
	Cost        int64        `gorm:"column:cost"`
	Name        string       `gorm:"column:name"`
	Description string       `gorm:"column:description"`
	Barcode     string       `gorm:"column:barcode"`
	ExpiredAt   sql.NullTime `gorm:"column:expired_at"`
	Size        string       `gorm:"column:size"`
}

func (Product) TableName() string {
	return "product"
}

//todo
//타입 설정, json/gorm 추가, time 타입, size 타입
