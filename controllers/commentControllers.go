package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetComments(c *gin.Context)  {
	db := database.GetDB()
	
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment := []models.Comment{}

	// userID untuk mengambil photo berdasarkan ID si user login
	err := db.Where("photo_id = ?", photoId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}
	err := db.Where("id = ?", commentId).Where("photo_id = ?", photoId).Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func UpdateComment(c *gin.Context)  {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeleteCommentById(c *gin.Context)  {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	
	Comment := models.Comment{}
	err := db.Where("user_id = ?", userID).Where("id = ?", commentId).Where("photo_id = ?", photoId).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Deleted",
	})
}