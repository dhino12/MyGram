package routes

import (
	"mygram/middlewares"
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorization(), controllers.GetPhotoById)
		photoRouter.POST("/", controllers.PostPhoto)

		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhotoById)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/:photoId", controllers.GetComments)
		commentRouter.GET("/:photoId/:commentId", middlewares.PhotoAuthorization(), controllers.GetCommentById)
		commentRouter.POST("/:photoId", middlewares.PhotoAuthorization(), controllers.PostComment)

		commentRouter.PUT("/:photoId/:commentId", middlewares.PhotoAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:photoId/:commentId", middlewares.PhotoAuthorization(), controllers.DeleteCommentById)
	}

	sosmedRouter := r.Group("/comments")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.GET("/", controllers.GetSocialMedia)
		sosmedRouter.GET("/:sosmedId", middlewares.PhotoAuthorization(), controllers.GetSocialMediaById)
		sosmedRouter.POST("/", middlewares.PhotoAuthorization(), controllers.PostSocialMedia)

		sosmedRouter.PUT("/:sosmedId", middlewares.PhotoAuthorization(), controllers.UpdateSocialMedia)
		sosmedRouter.DELETE("/:sosmedId", middlewares.PhotoAuthorization(), controllers.DeleteSocialMediaById)
	}
	return r
}