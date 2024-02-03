package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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

const (
	testPhoneNumber = "01011112222"
	testPassword    = "test-password"
)

var jwt = config.JWT{
	SecretKey:     "test-jwt-secret",
	ExpiredMinute: 1,
}

func TestUserHandler_LoginPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request.URL.Query().Add("message", "your_message")

	uh := handler.NewUserHandler(nil, nil, false)
	uh.LoginPage(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_Login(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	req.PostForm = map[string][]string{
		"phone_number": {testPhoneNumber},
		"password":     {testPassword},
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
	hashPW, _ := uh.Auth.MakeHashPassword(testPassword)

	repo.EXPECT().GetUser(context.Background(), testPhoneNumber).Return(domain.User{
		ID:          0,
		PhoneNumber: testPhoneNumber,
		Password:    hashPW,
	}, nil)

	uh.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)

	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	queryValues, _ := url.ParseQuery(u.RawQuery)
	assert.Contains(t, queryValues.Get("message"), "")
}

func TestUserHandler_Login_IncorrectPassword(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	req.PostForm = map[string][]string{
		"phone_number": {testPhoneNumber},
		"password":     {"test-incorrect-password"},
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)

	ua := auth.NewUserAuth(jwt)
	uh := handler.NewUserHandler(repo, ua, false)
	hashPW, err := uh.Auth.MakeHashPassword(testPassword)
	if err != nil {
		fmt.Println(err)
	}
	repo.EXPECT().GetUser(context.Background(), testPhoneNumber).Return(domain.User{
		ID:          0,
		PhoneNumber: testPhoneNumber,
		Password:    hashPW,
	}, nil)

	uh.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)

	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	queryValues, _ := url.ParseQuery(u.RawQuery)
	assert.Contains(t, queryValues.Get("message"), "Unauthorized")
}

func TestUserHandler_Logout(t *testing.T) {

}

func TestUserHandler_Register(t *testing.T) {
	// req 생성
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", nil)
	req.PostForm = map[string][]string{
		"phone_number":     {testPhoneNumber},
		"password":         {testPassword},
		"password_confirm": {testPassword},
	}
	// gin test context 생성
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// mock repo 생성
	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)
	repo.EXPECT().GetUser(context.Background(), testPhoneNumber).Return(domain.User{}, nil)

	// handler 생성
	ua := auth.NewUserAuth(jwt)
	uh := handler.NewUserHandler(repo, ua, false)

	hashPW, _ := uh.Auth.MakeHashPassword(testPassword)
	repo.EXPECT().CreateUser(context.Background(), &domain.User{
		PhoneNumber: testPhoneNumber,
		Password:    hashPW,
	}).Return(nil)

	uh.Register(c)

	assert.Equal(t, http.StatusOK, w.Code)

	// 현재 코드 기준 query param 메세지로 판단
	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}
	queryValues, _ := url.ParseQuery(u.RawQuery)
	assert.Contains(t, queryValues.Get("message"), "SuccessCreateAccount")
}

func TestUserHandler_Register_INVALID_PHONE_NUMBER(t *testing.T) {

}

func TestUserHandler_Register_ALREADY_EXIST_PHONE_NUMBER(t *testing.T) {

}

func TestUserHandler_Register_DOES_NOT_MATCH_CONFIRM_PASSWORD(t *testing.T) {

}
