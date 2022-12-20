package controllers

import (
	"net/http"
	"todoapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateBlogPostInput struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Source      string `json:"source"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type UpdateBlogPostInput struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Source      string `json:"source"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

// GET /tasks
// Get all tasks
func FindBlogPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post
	db.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// POST /tasks
// Create new task
func CreateBlogPost(c *gin.Context) {
	// Validate input
	var input CreateBlogPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	blogPost := models.Post{Author: input.Author, Title: input.Title, Description: input.Description, URL: input.URL, Source: input.Source, Category: input.Category, Language: input.Language, Country: input.Country}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&blogPost)

	c.JSON(http.StatusOK, gin.H{"data": blogPost})
}

// GET /posts/:id
// Find a blogpost
func FindPost(c *gin.Context) { // Get model if exist
	var post models.Post

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /posts/:id
// Update a blog post
func UpdatePost(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}

	// Validate input
	var input UpdateBlogPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Post
	updatedInput.Author = input.Author
	updatedInput.Title = input.Title
	updatedInput.Description = input.Description
	updatedInput.URL = input.URL
	updatedInput.Source = input.Source
	updatedInput.Category = input.Category
	updatedInput.Language = input.Language
	updatedInput.Country = input.Country

	db.Model(&post).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var blogPost models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&blogPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}

	db.Delete(&blogPost)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
