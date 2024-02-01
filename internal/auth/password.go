package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func (ua *UserAuth) MakeHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (ua *UserAuth) CheckPasswordHash(hashVal, userPw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw)) == nil
}
