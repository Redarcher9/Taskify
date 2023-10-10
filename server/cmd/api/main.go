package main

import (
	"taskify/api/routes"
	"taskify/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env := bootstrap.NewEnv()
	db := bootstrap.SetupDB(*env)
	routes.SetupRoutes(*env, db, r)
	r.Run() // listen and serve on 0.0.0.0:8080
	db.Close()
}
