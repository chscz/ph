package domain

type User struct {
	ID          int    `gorm:"column:id;primaryKey"`
	PhoneNumber string `gorm:"column:phone_number"`
	Password    string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}
