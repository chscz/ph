package router

import (
	"github.com/gin-gonic/gin"
	"payhere/internal/handler"
)

func InitGin(uh *handler.UserHandler, mh *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(AuthMiddleware(uh.Auth.JWTSecretKey))

	r.GET("/", mh.Home)

	r.GET("/login", uh.LoginPage)
	r.POST("/login", uh.Login)
	r.POST("/logout", uh.Logout)

	r.GET("/register", uh.RegisterPage)
	r.POST("/register", uh.Register)

	productGroup := r.Group("/product")
	{
		productGroup.GET("/create", mh.CreateProductPage)
		productGroup.POST("/create", mh.CreateProduct)

		productGroup.GET("/update/:id", mh.UpdateProductPage)
		productGroup.POST("/update/:id", mh.UpdateProduct)

		productGroup.GET("/delete/:id", mh.DeleteProduct)

		productGroup.GET("/detail/:id", mh.GetProductDetail)
		productGroup.GET("/search", mh.SearchProduct)
	}

	return r
}
