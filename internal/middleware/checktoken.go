package middleware

import (
	"net/http"

	"server/internal/util"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := util.VerifyAccessToken(cookie.Value, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)

		c.Next()
	}
}
