package routes

import (
	"github.com/Proudprogamer/goAuth/http/handlers"
	"github.com/Proudprogamer/goAuth/middleware"
	"github.com/gin-gonic/gin"
)


func SetUpRoutes(router *gin.Engine, handler *handlers.Handler) {

	router.GET("/home", handler.Home)
	router.POST("/sign-up", handler.SignUp)
	router.POST("/sign-in", handler.SignIn)

	protected := router.Group("/api")

	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handler.GetProfile)
	}

}