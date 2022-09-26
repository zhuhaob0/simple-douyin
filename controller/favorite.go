package controller

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	// token := c.PostForm("token")
	// 此处是POST请求，token不应该附在URL上，可能前端出错
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	logic.FavoriteAction(token, c)
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist!"})
		return
	}
	logic.FavoriteList(c)
}
