package routers

import (
	"github.com/gin-gonic/gin"
	"onycom/controllers"
	"onycom/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/signup", controllers.SignUp)
		userRoutes.POST("/signin", controllers.SignIn)
	}

	postRoutes := r.Group("/posts")
	{

		postRoutes.Use(middlewares.AuthMiddleware())
		{
			postRoutes.GET("", controllers.GetPosts)
			postRoutes.POST("/", controllers.CreatePost)
			postRoutes.GET("/:id", controllers.GetPost)
			postRoutes.PUT("/:id", controllers.UpdatePost)
			postRoutes.DELETE("/:id", controllers.DeletePost)
		}

	}

	return r
}
