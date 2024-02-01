package config

import "time"

type Config struct {
	MySQL MySQL
	JWT   JWT
}

func LoadConfig() (Config, error) {
	return Config{
		MySQL: MySQL{
			UserName: "root",
			Password: "1111",
			Host:     "localhost",
			Port:     "3306",
			Schema:   "ph",
		},
		JWT: JWT{
			SecretKey:   "jwtSecretKey",
			ExpiredTime: 10 * time.Minute,
			//ExpiredTime: 30 * time.Second,
		},
	}, nil
}
