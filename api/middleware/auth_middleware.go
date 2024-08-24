package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CORS() gin.HandlerFunc
}

type authMiddleware struct{}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (am *authMiddleware) CORS() gin.HandlerFunc {
	return cors.New(cors.Config{

		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
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
}
