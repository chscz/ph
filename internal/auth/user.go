package auth

import (
	"payhere/internal/config"
)

type UserAuth struct {
	JWTSecretKey     string
	JWTExpiredMinute int
}

func NewUserAuth(jwt config.JWT) *UserAuth {
	return &UserAuth{
		JWTSecretKey:     jwt.SecretKey,
		JWTExpiredMinute: jwt.ExpiredMinute,
	}
}
