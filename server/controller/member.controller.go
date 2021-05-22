package controller

import (
	"log"
	"morales-backend-1/database"
	"morales-backend-1/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController interface {
	GetMembers(c *gin.Context)
}

type memberController struct {
}

func MemberHandler() MemberController {
	return &memberController{}
}

func (controller memberController) GetMembers(c *gin.Context) {
	var members []model.Member

	if err := database.Conn.Raw("select Email, InsertDate from User").Find(&members).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	log.Println(members)

	c.JSON(http.StatusOK, gin.H{"data": members})
}
