package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sdlManager-mysql/model"
	"sdlManager-mysql/service"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}

	c.JSON(http.StatusOK, service.Login(user))
}

func GetUser(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	c.JSON(http.StatusOK, service.GetUser(name, currPage, pageSize, startTime, endTime))
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}

	c.JSON(http.StatusOK, service.CreateUser(user))

}

func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}

	c.JSON(http.StatusOK, service.UpdateUser(user))

}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, service.DeleteUser(id))
}
