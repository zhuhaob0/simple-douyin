package controller

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if _, exist := dao.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	logic.Publish(token, c)
}

// PublishList
// 用户发布的作品信息
func PublishList(c *gin.Context) {
	// 获取token
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	logic.PublishList(c)
}
