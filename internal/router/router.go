package router

import (
	"github.com/gin-gonic/gin"
	"payhere/internal/handler"
)

func InitGin(uh *handler.UserHandler, mh *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(AuthMiddleware(uh.Auth.JWTSecretKey))

	r.GET("/login", uh.LoginPage)
	r.POST("/login", uh.Login)
	r.POST("/logout", uh.Logout)

	r.GET("/register", uh.RegisterPage)
	r.POST("/register", uh.Register)

	r.GET("/", mh.Home)
	r.GET("/product/create", mh.CreateProductPage)
	r.POST("/product/create", mh.CreateProduct)

	r.GET("/product/detail/:id", mh.GetProductDetail)

	r.GET("/product/update/:id", mh.UpdateProductPage)
	r.POST("/product/update/:id", mh.UpdateProduct)

	r.GET("/product/delete/:id", mh.DeleteProduct)

	r.GET("/product/search", mh.SearchProduct)

	return r
}
