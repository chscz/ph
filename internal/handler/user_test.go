package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chscz/ph/internal/auth"
	"github.com/chscz/ph/internal/config"
	"github.com/chscz/ph/internal/domain"
	"github.com/chscz/ph/internal/handler"
	"github.com/chscz/ph/internal/handler/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_LoginPage(t *testing.T) {
	//uh := handler.NewUserHandler(nil, nil, false)
	//mock.New
}

func TestUserHandler_Login(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	req.PostForm = map[string][]string{
		"phone_number": {"01011112222"},
		"password":     {"test-password"},
	}

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)

	uh := handler.NewUserHandler(repo, nil, false)
	repo.EXPECT().GetUser(context.Background(), "01011112222").Return(domain.User{
		ID:          0,
		PhoneNumber: "01011112222",
		Password:    "test-password",
	}, nil)

	uh.Login(c)
}

func TestUserHandler_Login_IncorrectPassword(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	req.PostForm = map[string][]string{
		"phone_number": {"01011112222"},
		"password":     {"test-password"},
	}

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)

	jwt := config.JWT{
		SecretKey:     "test-jwt-secret",
		ExpiredMinute: 1,
	}

	ua := auth.NewUserAuth(jwt)
	uh := handler.NewUserHandler(repo, ua, false)
	hashPW, err := uh.Auth.MakeHashPassword("test-password")
	if err != nil {
		fmt.Println(err)
	}
	repo.EXPECT().GetUser(context.Background(), "01011112222").Return(domain.User{
		ID:          0,
		PhoneNumber: "01011112222",
		Password:    hashPW,
	}, nil)

	uh.Login(c)

	assert.Equal(t, http.StatusFound, w.Code)

	//u, err := url.Parse(w.Header().Get("Location"))
	//if err != nil {
	//	t.Fatalf("Failed to parse URL: %v", err)
	//}

	//queryValues, _ := url.ParseQuery(u.RawQuery)
	//assert.Contains(t, queryValues.Get("message"), "Unauthorized--")
}

func TestUserHandler_Logout(t *testing.T) {

}

func TestUserHandler_RegisterPage(t *testing.T) {

}

func TestUserHandler_Register(t *testing.T) {

}
