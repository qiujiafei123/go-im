package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-im/tools"
)

type FormCheckAuth struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

type FormRegister struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

type FormLogin struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

func Login(c *gin.Context) {
	var formLogin FormLogin
	if err := c.ShouldBindBodyWith(&formLogin, binding.JSON); err != nil {
		tools.FailWithMsg(c, err.Error())
	}

}


func CheckAuth(c *gin.Context) {
	var formCheckAuth FormCheckAuth
	if err := c.ShouldBindBodyWith(&formCheckAuth, binding.JSON); err != nil {
		tools.FailWithMsg(c, err.Error())
	}

	token := formCheckAuth.AuthToken
}


func Register(c *gin.Context) {
	var formRegister FormRegister
	if err := c.ShouldBindBodyWith(&formRegister, binding.JSON); err != nil {
		tools.FailWithMsg(c, err.Error())
	}
}


