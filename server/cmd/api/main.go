package main

import (
	"fmt"
	"taskify/api/routes"
	"taskify/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env := bootstrap.NewEnv()
	db := bootstrap.SetupDB(*env)
	error := db.Ping()
	routes.SetupRoutes(*env, db, r)
	if error != nil {
		fmt.Println("Db is working")
	}

	r.Run() // listen and serve on 0.0.0.0:8080
	db.Close()
}
