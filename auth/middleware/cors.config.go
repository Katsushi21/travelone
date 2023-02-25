package middleware

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	origin := os.Getenv("REQUEST_ORIGIN")

	c := cors.New(cors.Config{
		AllowOrigins: []string{origin},
		AllowMethods: []string{
			"POST",
			"GET",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})

	return c
}
