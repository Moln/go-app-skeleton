package controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/module/app/entity"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

type LoginParam struct {
	Username string
	Password string
}

func (c *AuthController) Login(ctx *gin.Context) {

	loginParam := &LoginParam{}
	ctx.Bind(loginParam)
	var user = entity.User{
		Username: loginParam.Username,
	}
	c.db.Table(entity.UserTable).First(&user)

	if user.PasswordVerify(loginParam.Password) {

	}
}

func (c *AuthController) Logout(ctx *gin.Context) {

}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db,
	}
}
