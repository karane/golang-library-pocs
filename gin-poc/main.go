package main

import (
	"gin-poc/middleware"
	"gin-poc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.HeaderMiddleware())
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api/v1")
	routes.RegisterUserRoutes(api.Group("/users"))

	r.Run(":8080")
}
