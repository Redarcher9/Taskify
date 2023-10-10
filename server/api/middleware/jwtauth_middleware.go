package middleware

import (
	"net/http"
	"taskify/utils"

	"github.com/gin-gonic/gin"
)

/*
Description: JWT authentication middleware for protected routes

Params: secret: HMAC key for encryption

Returns: gin  handlerfunc for middleware
*/
func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Auth_Token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			c.Abort()
			return
		}
		err = utils.ValidateToken(c, cookie, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
