package mysql

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/chscz/ph/internal/config"
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMYSQL(cfg config.MySQL) (*gorm.DB, error) {
	mysqlCfg := &mysql.Config{
		User:   cfg.UserName,
		Passwd: cfg.Password,
		Net:    "tcp",
		Addr:   net.JoinHostPort(cfg.Host, cfg.Port),
		DBName: cfg.DB,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		Loc:                  time.UTC,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var db *gorm.DB
	var err error
	retryCount := 20
	for i := 0; i < retryCount; i++ {
		dsn := mysqlCfg.FormatDSN()
		db, err = gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to MySQL. Retrying (%d/%d)...\n", i+1, retryCount)
		time.Sleep(time.Second * 5)
		if i == retryCount-1 {
			panic(err)
		}
	}

	conn, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("gorm db failed : %w", err)
	}
	conn.SetMaxIdleConns(3)
	conn.SetMaxOpenConns(3)
	return db, nil
}
