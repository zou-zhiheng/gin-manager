package router

import (
	"github.com/gin-gonic/gin"
	"sdlManager-mysql/controller"
	"sdlManager-mysql/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()

	//跨域
	engine.Use(middleware.Cors())
	engine.POST("login", controller.Login)
	//权限
	engine.Use(middleware.JWTAuth(), middleware.ApiAuth())
	//用户
	UserRouter(engine)
	//角色
	RoleRouter(engine)
	//api
	ApiRouter(engine)

	return engine
}
