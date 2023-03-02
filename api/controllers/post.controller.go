package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theifedayo/go-posts/api/models"
	"github.com/theifedayo/go-posts/config"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Content: body.Content}
	result := config.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "post created successfully",
		"data":    post,
	})

}

func PostsList(c *gin.Context) {

	var posts []models.Post
	config.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "posts retrieved successfully",
		"data":    posts,
	})
}
