package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sdlManager-mysql/model"
	"sdlManager-mysql/service"
)

func GetRole(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	c.JSON(http.StatusOK, service.GetRole(name, currPage, pageSize, startTime, endTime))
}

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}

	c.JSON(http.StatusOK, service.CreateRole(role))

}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}

	c.JSON(http.StatusOK, service.UpdateRole(role))

}

func DeleteRole(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, service.DeleteRole(id))
}
