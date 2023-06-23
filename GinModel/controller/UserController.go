package controller

import (
	"clanrece.com/EchoPong/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginIn(c *gin.Context) {
	var userInfo model.User
	_ = c.ShouldBindJSON(&userInfo)
	err := model.Login(&userInfo)
	if err != nil {
		c.String(http.StatusInternalServerError, "{\"status\": \"500\"}")
		return
	}
	c.String(http.StatusOK, "{\"status\": \"200\"}")
}

func ListUsers(c *gin.Context) {
	pageNo := c.GetInt("pageNo")
	pageSize := c.GetInt("pageSize")
	users := model.ListUsers(pageNo, pageSize)
	c.JSON(http.StatusOK, users)
}

func UserDetail(c *gin.Context) {
	value := c.Query("userName")
	userInfo := model.GetUserDetailByName(value)
	c.JSON(http.StatusOK, userInfo)
}
