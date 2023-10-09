package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(env map[string]string, db *sql.DB, gin *gin.Engine) {
	Router := gin.Group("")
	NewPingRoute(env, db, Router)
	NewLoginRoute(env, db, Router)
	NewSignupRoute(env, db, Router)
}
