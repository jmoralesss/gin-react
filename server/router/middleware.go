package router

import (
	"log"
	"morales-backend-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BearerTokenMiddleware(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if _, err := jwtService.ValidateToken(token); err != nil {
			log.Println(err.Error())

			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			c.Abort()
		}

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
