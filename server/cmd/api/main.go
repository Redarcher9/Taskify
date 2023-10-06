package main

import (
	"fmt"
	"net/http"
	"taskify/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := bootstrap.SetupDB()
	error := db.Ping()
	if error != nil {
		fmt.Println("Db is working")
	}
	env := bootstrap.NewEnv()
	fmt.Println(env)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	db.Close()
}
