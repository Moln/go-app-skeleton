package controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/module/app/entity"
	"gorm.io/gorm"
	"net/http"
)

type AppController struct {
	db *gorm.DB
}

func (e *AppController) GetUsers(c *gin.Context) {
	var users []entity.User

	e.db.Table("admin_users").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (e *AppController) GetUser(c *gin.Context) {
	userId := c.Params.ByName("id")
	var user entity.User
	e.db.First(&user, userId)

	c.JSON(http.StatusOK, user)
}

func NewAppController(db *gorm.DB) *AppController {
	return &AppController{
		db,
	}
}
