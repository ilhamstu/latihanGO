package controllers

import (
	"errors"
	"net/http"
	"restapi/be/models"
	"restapi/be/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func FindPost(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data",
		"data":    posts,
	})
}

func StorePost(c *gin.Context) {
	var input validators.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Created",
		"data":    post,
	})
}

func FindPostByID(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail data Post by ID : " + c.Param("id"),
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input validators.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	if err := models.DB.Model(&post).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := models.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}
