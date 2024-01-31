package http

import (
	"github.com/gin-gonic/gin"
)

func InitGin(uh UserRoute, mh MenuRoute) *gin.Engine {
	r := gin.Default()

	// 로그인  jwt token 기반, 세션유지
	r.GET("/login", uh.login)
	//r.GET("/login", handler.Home)
	//r.GET("/login", handler.Tee{}.Home)
	// 가입  폰번호 유효성검사, 중복검사, db 저장
	r.GET("/register", uh.register)

	// 세션없을시 /login으로 redirect
	// 상품 리스트 crud 및 like/초성 검색 가능
	// cursor based pagination 10개씩
	// 리스트의 최상단에는 검색/신규 추가
	// 리스트 항목의 오른쪽에 수정/삭제 버튼
	r.GET("/", mh.home)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
