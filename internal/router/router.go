package router

import (
	"github.com/chscz/ph/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitGin(uh *handler.UserHandler, ph *handler.ProductHandler) *gin.Engine {
	r := gin.Default()

	// intellij debug 환경
	//r.LoadHTMLGlob("templates/*")
	// docker 환경
	r.LoadHTMLGlob("/app/templates/*")

	r.Use(AuthMiddleware(uh.Auth.JWTSecretKey))

	r.GET("/", ph.Home)

	r.GET("/login", uh.LoginPage)
	r.POST("/login", uh.Login)
	r.POST("/logout", uh.Logout)

	r.GET("/register", uh.RegisterPage)
	r.POST("/register", uh.Register)

	productGroup := r.Group("/product")
	{
		productGroup.GET("/create", ph.CreateProductPage)
		productGroup.POST("/create", ph.CreateProduct)

		productGroup.GET("/update/:id", ph.UpdateProductPage)
		productGroup.POST("/update/:id", ph.UpdateProduct)

		productGroup.GET("/delete/:id", ph.DeleteProduct)

		productGroup.GET("/detail/:id", ph.GetProductDetail)
		productGroup.GET("/search", ph.SearchProduct)
	}

	return r
}
