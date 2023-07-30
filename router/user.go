package router

import (
	"github.com/gin-gonic/gin"
	"sdlManager-mysql/controller"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.GET("get", controller.GetUser)
		user.POST("create", controller.CreateUser)
		user.PUT("update", controller.UpdateUser)
		user.DELETE("delete", controller.DeleteUser)
	}
}
