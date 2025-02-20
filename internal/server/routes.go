package server

import (
	"net/http"
	"store-management/internal/handlers"
	middlewares "store-management/internal/middleware"
	"store-management/internal/repositories"
	"store-management/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	storeRepo := repositories.NewStoreRepository(s.db.GetDB())
	storeService := services.NewStoreService(storeRepo)
	storeHandler := handlers.NewStoreHandler(storeService)

	r.POST("/stores", storeHandler.CreateStore)

	auth := r.Group("/")
	auth.Use(middlewares.Auth())
	{
		storeGroup := auth.Group("/stores")
		{
			storeGroup.GET("/:id", storeHandler.GetStoreByID)
			storeGroup.PUT("/:id", storeHandler.UpdateStore)
			storeGroup.DELETE("/:id", storeHandler.DeleteStore)
			storeGroup.GET("/", storeHandler.ListStores)
		}
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
