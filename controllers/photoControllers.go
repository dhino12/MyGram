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

func PostPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context)  {
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

func GetPhotos(c *gin.Context)  {
	db := database.GetDB()
	
	Photo := []models.Photo{}

	// userID untuk mengambil photo berdasarkan ID si user login
	err := db.Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo := models.Photo{}
	err := db.Where("id = ?", photoId).Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhotoById(c *gin.Context)  {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	
	Photo := models.Photo{}
	err := db.Where("user_id = ?", userID).Where("id = ?", photoId).Delete(&Photo).Error
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