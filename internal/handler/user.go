package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
}

func NewUserHandler(repo UserRepository) UserHandler {
	return UserHandler{repo: repo}
}

func (uh *UserHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
	})
}

func (uh *UserHandler) Register(c *gin.Context) {

}
