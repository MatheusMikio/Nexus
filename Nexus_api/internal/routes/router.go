package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	container := container.NewContainer(db)
	router := gin.Default()
	initRoutes(router, container)	
	router.Run(":8080")
}
