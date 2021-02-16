package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key`
	Name   string `json:"name"`
	Auther string `json:"auther"`
}

type CreateBookInput struct {
	Name   string `json:"name" binding:"required"`
	Auther string `json:"auther" binding:"required"`
}
type UpdateBookInput struct {
	Name   string `json:"name"`
	Auther string `json:"auther"`
}

func SetupModels() *gorm.DB {

	db, err := gorm.Open("mysql", "root:newpassword@tcp(127.0.0.1:3306)/febin?charset=utf8&parseTime=True")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	return db
}

// GET /posts
// Get all posts
func FindPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /posts
// Create new posts
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create post
	book := Book{Name: input.Name, Auther: input.Auther}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /posts/:id
// Find a post
func FindPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /post/:id
// Update a post
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func main() {
	r := gin.Default()
	db := SetupModels() // new
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.GET("/posts", FindPosts)
	r.POST("/posts", CreatePost)       // create
	r.GET("/posts/:id", FindPost)      // find by id
	r.PATCH("/posts/:id", UpdatePost)  // update by id
	r.DELETE("/posts/:id", DeletePost) // delete by id
	r.Run(":7000")
}
