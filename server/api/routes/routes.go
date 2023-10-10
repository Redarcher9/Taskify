package routes

import (
	"database/sql"
	"taskify/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(env map[string]string, db *sql.DB, gin *gin.Engine) {
	Router := gin.Group("")
	NewLoginRoute(env, db, Router)
	NewSignupRoute(env, db, Router)
	ProtectedRoute := gin.Group("")
	ProtectedRoute.Use(middleware.JwtAuthMiddleware(""))
	NewPingRoute(env, db, ProtectedRoute)
}
