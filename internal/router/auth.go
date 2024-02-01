package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"payhere/internal/auth"
)

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.URL.Path == "/login") || (c.Request.URL.Path == "/register") {
			c.Next()
			return
		}

		accessToken, err := c.Cookie("access-token")
		if err != nil {
			//todo
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		claims, err := auth.ValidateJWT(jwtKey, accessToken)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				c.SetCookie("access-token", "", -1, "/", "", false, true)
			}
			//todo
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
