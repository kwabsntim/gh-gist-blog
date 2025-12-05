package handlers

import (
	"ghgist-blog/middleware"
	"ghgist-blog/services"

	"github.com/gin-gonic/gin"
)

// ServiceContainer holds all service interfaces

type ServiceContainer struct {
	RegisterService services.RegisterInterface
	LoginService    services.LoginInterface
	FetchService    services.FetchWritersInterface
}

func RouteSetup(services *ServiceContainer) *gin.Engine {

	handlers := NewHandlers(services)

	//gin router
	router := gin.Default()

	//public routes
	public := router.Group("/api/public")
	{
		public.GET("/writers", handlers.FetchAllWriters)
		public.POST("/auth/login", handlers.Login)
	}
	//Protected based routes
	protected := router.Group("api/admin")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/CreateArticles")
	}

	//Role based routes
	adminRoute := router.Group("/api/admin")
	adminRoute.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware())
	{
		adminRoute.POST("/auth/register", handlers.Register)
	}
	return router

}
