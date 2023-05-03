package middlewares

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		fmt.Println(userData)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":  "Data Not Found",
				"message": "data doesn't exist",
			})

			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorization",
				"message": "you are not allowed to access this data",
			})

			return
		}

		ctx.Next()
	}
}