package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theifedayo/go-posts/config"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
