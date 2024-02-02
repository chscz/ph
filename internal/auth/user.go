package auth

import (
	"github.com/chscz/ph/internal/config"
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
