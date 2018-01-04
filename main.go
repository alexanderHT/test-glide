package main

import (
	"controllers/jsc"
	"vendor/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	JscLogin()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
