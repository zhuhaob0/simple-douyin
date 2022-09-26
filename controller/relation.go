package controller

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	dao.Response
	UserList []dao.User `json:"user_list"`
}

// RelationAction 关注或取消关注
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	if user, exist := dao.UsersLoginInfo[token]; exist {
		actionType := c.Query("action_type")
		if actionType == "1" { //关注
			logic.Follow(user.ID, c)
		} else { // 取消关注
			logic.CancelFollow(user.ID, c)
		}
		c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; exist {
		userID, _ := strconv.Atoi(c.Query("user_id"))
		c.JSON(http.StatusOK, UserListResponse{
			Response: dao.Response{
				StatusCode: 0,
			},
			UserList: logic.FollowList(int64(userID)),
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"}})
	}
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := dao.UsersLoginInfo[token]; exist {
		userID, _ := strconv.Atoi(c.Query("user_id"))
		c.JSON(http.StatusOK, UserListResponse{
			Response: dao.Response{
				StatusCode: 0,
			},
			UserList: logic.FollowerList(int64(userID)),
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"}})
	}
}
