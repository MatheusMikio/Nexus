package middlewares

import (
	"net/http"
	"time"

	"github.com/MatheusMikio/Nexus/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:5173",
	}

	if configuredOrigins := config.GetCORSAllowedOrigins(); len(configuredOrigins) > 0 {
		allowedOrigins = configuredOrigins
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
