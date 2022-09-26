package controller

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction 发表或删除评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if user, exist := dao.UsersLoginInfo[token]; exist {
		if actionType == "1" { //评论
			logic.Comment(&user, c)
		} else { // 删除评论
			logic.DeleteComment(c)
		}
	} else {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList 返回评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; exist {
		comments := logic.CommentList(c)
		c.JSON(http.StatusOK, dao.CommentListResponse{
			Response:    dao.Response{StatusCode: 0},
			CommentList: comments,
		})
	} else {
		c.JSON(http.StatusOK, dao.CommentListResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
