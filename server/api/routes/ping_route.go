package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewPingRoute(env map[string]string, db *sql.DB, r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
}
