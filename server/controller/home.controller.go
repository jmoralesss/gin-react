package controller

import (
	"fmt"
	"morales-backend-1/database"
	"morales-backend-1/model"
	"morales-backend-1/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HomeController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type homeController struct {
	jWtService service.JWTService
}

func HomeHandler(jWtService service.JWTService) HomeController {
	return &homeController{
		jWtService: jWtService,
	}
}

func (controller homeController) Login(ctx *gin.Context) {
	var input model.LoginAttempt

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User

	// check if there is a user with the provided email
	if err := database.Conn.Where(model.User{Email: input.Email}).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var token string

	if service.AuthenticateAgainstHash(input.Password, user.Password) {
		token = controller.jWtService.GenerateToken(user.Email, true)
	}

	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}

func (controller homeController) Register(c *gin.Context) {
	var input model.RegistrationAttempt

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if provided email is already associated with a member
	var count int64
	if err := database.Conn.Model(&model.User{}).Select("ID").Where(fmt.Sprintf("%s = ?", "Email"), input.Email).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error occured processing your request"})
		return
	}

	user := model.User{Email: input.Email, InsertDate: time.Now()}

	// create hashed password from provided password
	if hashedPassword, err := service.HashPassword(input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user.Password = hashedPassword
	}

	// create new user record
	if err := database.Conn.Where(model.User{Email: input.Email}).FirstOrCreate(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
