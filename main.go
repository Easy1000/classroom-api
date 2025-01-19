package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	initDB()

	// Seed the database with initial data
	seedDatabase()

	// Create a Gin router
	r := gin.Default()

	// Define API routes
	r.GET("/students", getStudents)
	r.GET("/classes", getClasses)
	r.GET("/enrollments", getEnrollments)
	r.POST("/comments", createComment)
	r.GET("/comments", getComments)
	r.GET("/comments/:id", getCommentByID)
	r.DELETE("/comments/:id", deleteComment)

	// Start the server
	r.Run(":8080")
}

func getStudents(c *gin.Context) {
	var students []Student
	if result := db.Find(&students); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func getClasses(c *gin.Context) {
	var classes []Class
	if result := db.Find(&classes); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func getEnrollments(c *gin.Context) {
	var enrollments []Enrollment
	if result := db.Find(&enrollments); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollments)
}

func createComment(c *gin.Context) {
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&comment)
	c.JSON(http.StatusOK, comment)
}

func getComments(c *gin.Context) {
	var comments []Comment
	db.Find(&comments)
	c.JSON(http.StatusOK, comments)
}

func getCommentByID(c *gin.Context) {
	id := c.Param("id")
	var comment Comment
	if err := db.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func deleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Comment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
