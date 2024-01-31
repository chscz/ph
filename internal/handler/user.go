package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/internal/domain"
	"regexp"
)

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, phoneNumber string) (domain.User, error)
}

func NewUserHandler(repo UserRepository) UserHandler {
	return UserHandler{repo: repo}
}

func (uh *UserHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	ctx := context.Background()
	pn := c.PostForm("phone_number")
	pw := c.PostForm("password")

	user, err := uh.repo.GetUser(ctx, pn)
	if err != nil {
		//todo
	}

	if pw != user.Password {
		//todo
	}
	//todo
	c.Redirect(http.StatusFound, "/foo")
}

func (uh *UserHandler) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"title": "register page",
	})
}

func (uh *UserHandler) Register(c *gin.Context) {
	ctx := context.Background()
	pn, ok := c.GetPostForm("phone_number")
	if !ok {
		//todo
	}
	pw, ok := c.GetPostForm("password")
	if !ok {
		//todo
	}

	if !isValidPhoneNumber(pn) {
		return
	}

	user := &domain.User{
		ID:          0,
		PhoneNumber: pn,
		Password:    pw,
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
