package router

import (
	"github.com/gin-gonic/gin"
	"sdlManager-mysql/controller"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.GET("get", controller.GetApi)
		api.POST("create", controller.CreateApi)
		api.PUT("update", controller.UpdateApi)
		api.DELETE("delete", controller.DeleteApi)
	}
}
