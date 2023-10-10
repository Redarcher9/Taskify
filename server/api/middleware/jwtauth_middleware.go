package middleware

import (
	"net/http"
	"taskify/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateToken(c, token)
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
