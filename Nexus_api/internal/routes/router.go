package routes

import (
	"github.com/MatheusMikio/Nexus/internal/container"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	container := container.NewContainer(db)
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	initRoutes(router, container)
	router.Run(":8080")
}
