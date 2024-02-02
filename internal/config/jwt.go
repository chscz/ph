package config

type JWT struct {
	SecretKey     string `env:"SECRET_KEY"`
	ExpiredMinute int    `env:"EXPIRED_MINUTE"`
}
