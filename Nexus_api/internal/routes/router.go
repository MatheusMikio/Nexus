package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	// container := container.NewContainer(db)
	router.Run(":8080")
}
