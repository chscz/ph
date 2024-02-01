package config

import "time"

type JWT struct {
	SecretKey   string
	ExpiredTime time.Duration
}
