package mysql

import "gorm.io/gorm"

type UserRepo struct {
	DB *gorm.DB
}

type ProductRepo struct {
	DB *gorm.DB
}
