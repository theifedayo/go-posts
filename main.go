package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theifedayo/go-posts/api/controllers"
	"github.com/theifedayo/go-posts/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsList)
	r.GET("/posts/:id", controllers.PostsDetail)

	r.Run()

}
