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

func PostsDetail(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	config.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "post retrieved successfully",
		"data":    post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)
	var post models.Post
	config.DB.First(&post, id)

	config.DB.Model(&post).Updates(models.Post{Title: body.Title, Content: body.Content})

	c.JSON(200, gin.H{
		"message": "post retrieved successfully",
		"data":    post,
	})
}
