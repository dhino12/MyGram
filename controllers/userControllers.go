package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_,_ = db, contentType
	UserData := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&UserData)
	} else {
		ctx.ShouldBind(&UserData)
	}
	
	err := db.Debug().Create(&UserData).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": UserData.ID,
		"email": UserData.Email,
		"username": UserData.Username,
		"age": UserData.Age,
	})
	return
}

func UserLogin(c *gin.Context)  {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H {
		"token": token,
	})
}