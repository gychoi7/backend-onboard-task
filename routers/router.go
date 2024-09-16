package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"onycom/controllers"
	"onycom/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
