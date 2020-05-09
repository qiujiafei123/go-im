package router

import (
	"github.com/gin-gonic/gin"
	"go-im/api/handler"
	"go-im/tools"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		tools.FailWithMsg(c, "路由不存在")
	})
	userRouter(r)
	return r
}


func userRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/login", handler.Login)
}