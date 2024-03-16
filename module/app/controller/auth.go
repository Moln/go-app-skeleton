package controller

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-demo/common/response"
	"go-demo/module/app/entity"
	"gorm.io/gorm"
	"net/http"
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
	db := c.db.Table(entity.UserTable).First(&user, user)

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {

		} else {
			ctx.Error(db.Error)
			return
		}
	}

	if !user.PasswordVerify(loginParam.Password) {
		response.Error(ctx, "账号或密码错误")
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("user", user.Id)
	sess.Save()
	//sess := ctx.MustGet("session").(string)

	ctx.Status(http.StatusNoContent)
}

func (c *AuthController) Logout(ctx *gin.Context) {

	sess := sessions.Default(ctx)
	sess.Clear()
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	sess.Save()

	ctx.Status(http.StatusNoContent)
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db,
	}
}
