package domain

import "database/sql"

type Menu struct {
	ID          int
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Category    string
	Price       int
	Cost        int
	Name        string
	Description string
	Barcode     string
	ExpiredAt   sql.NullTime
	Size        string
}

//todo
//타입 설정, json/gorm 추가, time 타입, size 타입
