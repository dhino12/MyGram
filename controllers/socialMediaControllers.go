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

func PostSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context)  {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))
	// userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	// SocialMedia.UserID = userID
	// SocialMedia.ID = uint(sosmedId)

	err := db.Model(&SocialMedia).Where("id = ?", sosmedId).Updates(models.SocialMedia{
		Name: SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func GetSocialMedia(c *gin.Context)  {
	db := database.GetDB()
	
	SocialMedia := []models.SocialMedia{}

	// userID untuk mengambil photo berdasarkan ID si user login
	err := db.Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	SocialMedia := models.SocialMedia{}
	err := db.Where("id = ?", sosmedId).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMediaById(c *gin.Context)  {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))
	userID := uint(userData["id"].(float64))
	
	SocialMedia := models.SocialMedia{}
	err := db.Where("user_id = ?", userID).Where("id = ?", sosmedId).Delete(&SocialMedia).Error
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