package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"payhere/internal/auth"
	"payhere/internal/domain"
	"regexp"
)

type UserHandler struct {
	repo UserRepository
	Auth *auth.UserAuth
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, phoneNumber string) (domain.User, error)
}

func NewUserHandler(repo UserRepository, auth *auth.UserAuth) *UserHandler {
	return &UserHandler{repo: repo, Auth: auth}
}

func (uh *UserHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_login.tmpl", gin.H{
		"title":   "로그인",
		"message": c.Query("message"),
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	ctx := context.Background()
	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")

	// 계정 조회
	user, err := uh.repo.GetUser(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Redirect(http.StatusFound, "/login?message=NotExistUser")
			return
		}
		c.Redirect(http.StatusFound, "/login?message=InternalError")
		return
	}
	// 비밀번호 검사
	if !uh.Auth.CheckPasswordHash(user.Password, password) {
		c.Redirect(http.StatusFound, "/login?message=Unauthorized")
		return
	}

	// 토큰 발행
	accessToken, err := uh.Auth.CreateJWT(user.PhoneNumber)
	if err != nil {
		//todo
	}

	cookie, err := c.Cookie("access-token")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie(
			"access-token",
			accessToken,
			3600,
			"/",
			"localhost",
			false,
			true,
		)
	}
	_ = cookie

	c.Redirect(http.StatusFound, "/")
}

func (uh *UserHandler) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_register.tmpl", gin.H{
		"title":   "회원가입",
		"message": c.Query("message"),
	})
}

func (uh *UserHandler) Register(c *gin.Context) {
	ctx := context.Background()
	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("password_confirm")

	// 휴대폰 번호 유효성 검사
	if !isValidPhoneNumber(phoneNumber) {
		c.Redirect(http.StatusFound, "/register?message=InvalidPhoneNumber")
		return
	}
	// 비밀번호 검사
	if password != passwordConfirm {
		c.Redirect(http.StatusFound, "/register?message=PasswordsDoNotMatch")
		return
	}
	// 기존 유저 phone number 중복 여부 검사
	u, err := uh.repo.GetUser(ctx, phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.Redirect(http.StatusFound, "/register?message=InternalError")
		return
	}
	if u.PhoneNumber != "" {
		c.Redirect(http.StatusFound, "/register?message=AlreadyExistPhoneNumber")
		return
	}
	// 비밀번호 암호화
	hashPassword, err := uh.Auth.MakeHashPassword(password)
	if err != nil {
		//todo
	}

	user := &domain.User{
		PhoneNumber: phoneNumber,
		Password:    hashPassword,
	}

	if err := uh.repo.CreateUser(ctx, user); err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusFound, "/login")
}

func isValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(`^010\d{8}$`)
	return phoneRegex.MatchString(phoneNumber)
}
