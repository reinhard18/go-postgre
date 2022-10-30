package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reinhard18/go-postgre/initializers"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
