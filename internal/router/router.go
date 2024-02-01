package router

import (
	"github.com/gin-gonic/gin"
	"payhere/internal/handler"
)

func InitGin(uh *handler.UserHandler, mh *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(AuthMiddleware(uh.Auth.JWTSecretKey))

	// 로그인  jwt token 기반, 세션유지
	r.GET("/login", uh.LoginPage)
	r.POST("/login", uh.Login)
	// 가입  폰번호 유효성검사, 중복검사, db 저장
	r.GET("/register", uh.RegisterPage)
	r.POST("/register", uh.Register)

	// 세션없을시 /login 으로 redirect
	// 상품 리스트 crud 및 like/초성 검색 가능
	// cursor based pagination 10개씩
	// 리스트의 최상단에는 검색/신규 추가
	// 리스트 항목의 오른쪽에 수정/삭제 버튼
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
