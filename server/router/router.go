package router

import (
	"morales-backend-1/controller"
	"morales-backend-1/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
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

	return router
}
