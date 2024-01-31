package http

import (
	"github.com/gin-gonic/gin"
	"payhere/internal/handler"
)

type UserRoute struct {
	handler handler.UserHandler
}

func (uh *UserRoute) login(c *gin.Context) {
	return
}

func (uh *UserRoute) register(c *gin.Context) {
	return
}
