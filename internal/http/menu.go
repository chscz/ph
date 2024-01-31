package http

import (
	"github.com/gin-gonic/gin"
	"payhere/internal/handler"
)

type MenuRoute struct {
	handler handler.MenuHandler
}

func (mr *MenuRoute) home(c *gin.Context) {
	
	mr.handler.Home()
	return
}
