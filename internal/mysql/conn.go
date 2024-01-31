package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB() *sql.DB {
	connectionString := "username:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")
	return db
}
