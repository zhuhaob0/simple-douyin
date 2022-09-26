package controller

import (
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	logic.Register(username, password, c)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	logic.Login(username, password, c)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	logic.UserInfo(token, c)
}
