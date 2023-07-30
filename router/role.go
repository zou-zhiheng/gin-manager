package router

import (
	"github.com/gin-gonic/gin"
	"sdlManager-mysql/controller"
)

func RoleRouter(engine *gin.Engine) {

	role := engine.Group("role")
	{
		role.GET("get", controller.GetRole)
		role.POST("create", controller.CreateRole)
		role.PUT("update", controller.UpdateRole)
		role.DELETE("delete", controller.DeleteRole)
	}
}
