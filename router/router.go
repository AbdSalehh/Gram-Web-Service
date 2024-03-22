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
			photoRouter.GET("/all", controllers.PhotoGetAll)
			photoRouter.GET("/", controllers.PhotoGet)      
			photoRouter.POST("/", controllers.PhotoCreate)
			photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)
			photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
		}

		commentRouter := r.Group("/comments")
		{
			commentRouter.GET("/all", controllers.CommentGetAll)
			commentRouter.GET("/", controllers.CommentGet)
			commentRouter.POST("/", controllers.CommentCreate)
			commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.CommentUpdate)
			commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.CommentDelete)
		}

		socialMediaRouter := r.Group("/socialmedias")
		{
			socialMediaRouter.GET("/all", controllers.SocialMediaGetAll)
			socialMediaRouter.GET("/", controllers.SocialMediaGet)
			socialMediaRouter.POST("/", controllers.SocialMediaCreate)
			socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaUpdate)
			socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaDelete)
		}
	}

	return r
}

