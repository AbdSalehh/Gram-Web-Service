package router

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	middlewareRouter := r.Use(middlewares.Authentication())
	{
		middlewareRouter.PUT("/users", middlewares.UserAuthorization(), controllers.UpdateUser)
		middlewareRouter.DELETE("/users", middlewares.UserAuthorization(), controllers.DeleteUser)

		photoRouter := r.Group("/photos")
		{
			photoRouter.GET("/", controllers.PhotoGetAll)
			photoRouter.POST("/", controllers.PhotoCreate)
			photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)
			photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
		}

		commentRouter := r.Group("/comments")
		{
			commentRouter.GET("/", controllers.CommentGetAll)
			commentRouter.POST("/", controllers.CommentCreate)
			commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.CommentUpdate)
			commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.CommentDelete)
		}

		socialMediaRouter := r.Group("/socialmedias")
		{
			socialMediaRouter.GET("/", controllers.SocialMediaGetAll)
			socialMediaRouter.POST("/", controllers.SocialMediaCreate)
			socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaUpdate)
			socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaDelete)
		}
	}

	return r
}

