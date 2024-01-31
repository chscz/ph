package domain

import "database/sql"

type User struct {
	ID          int
	CreatedAt   sql.NullTime
	PhoneNumber string
	Password    string
	AccessedAt  sql.NullTime
}
