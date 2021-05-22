package main

import (
	"log"
	"morales-backend-1/controller"
	"morales-backend-1/database"
	"morales-backend-1/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	jwtService := service.JWTAuthService()

	homeController := controller.HomeHandler(jwtService)
	memberController := controller.MemberHandler()

	router.Use(CORSMiddleware())

	router.GET("/keep-alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": time.Now()})
	})

	router.POST("/login", homeController.Login)
	router.POST("/register", homeController.Register)

	authorized := router.Group("/member")
	authorized.Use(BearerTokenMiddleware(jwtService))
	{
		authorized.GET("/", memberController.GetMembers)
	}

	database.ConnectDatabase()

	router.RunTLS(":443", "./.cert/cert.pem", "./.cert/key.pem")
}

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
