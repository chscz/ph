package auth

import (
	"payhere/internal/config"
	"time"
)

type UserAuth struct {
	JWTSecretKey   string
	JWTExpiredTime time.Duration
}

func NewUserAuth(jwt config.JWT) *UserAuth {
	return &UserAuth{
		JWTSecretKey:   jwt.SecretKey,
		JWTExpiredTime: jwt.ExpiredTime,
	}
}
