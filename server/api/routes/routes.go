package routes

import (
	"database/sql"
	"taskify/api/middleware"

	"github.com/gin-gonic/gin"
)

/*
Description: Sets up routes in main file

Params: env - map of env variables, db - database instance, gin - gin engine

Returns: NA
*/

func SetupRoutes(env map[string]string, db *sql.DB, gin *gin.Engine) {
	Router := gin.Group("")
	NewLoginRoute(env, db, Router)
	NewSignupRoute(env, db, Router)
	ProtectedRoute := gin.Group("")
	ProtectedRoute.Use(middleware.JwtAuthMiddleware(env["HMAC_KEY"]))
	NewPingRoute(env, db, ProtectedRoute)
}
