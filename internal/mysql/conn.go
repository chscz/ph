package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"payhere/internal/config"
	"time"
)

func InitMYSQL(cfg config.MySQL) (*gorm.DB, error) {
	mysqlCfg := &mysql.Config{
		User:   cfg.UserName,
		Passwd: cfg.Password,
		Net:    "tcp",
		Addr:   net.JoinHostPort(cfg.Host, cfg.Port),
		DBName: cfg.Schema,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		Loc:                  time.UTC,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	dsn := mysqlCfg.FormatDSN()

	db, err := gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open failed : %w", err)
	}

	conn, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("gorm db failed : %w", err)
	}
	conn.SetMaxIdleConns(3)
	conn.SetMaxOpenConns(3)
	return db, nil
}
