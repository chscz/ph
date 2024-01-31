package handler

import "github.com/gin-gonic/gin"

type MenuHandler struct {
	repo MenuRepository
}

type MenuRepository interface {
}

func NewMenuHandler(repo MenuRepository) MenuHandler {
	return MenuHandler{repo: repo}
}

func (mh *MenuHandler) Home(c *gin.Context) {
	return
}

func (mh *MenuHandler) Create(c *gin.Context) {
	return
}

func (mh *MenuHandler) Update(c *gin.Context) {
	return
}

func (mh *MenuHandler) Delete(c *gin.Context) {
	return
}

func (mh *MenuHandler) Search(c *gin.Context) {
	return
}
