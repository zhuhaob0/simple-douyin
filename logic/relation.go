package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Follow(userID int64, c *gin.Context) {
	toUserID, _ := strconv.Atoi(c.Query("to_user_id"))
	if toUserID == 0 {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "无法关注admin"})
		return
	} else if userID == int64(toUserID) {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "无法关注自己"})
		return
	}
	//println("Follow", userID, toUserID)
	err := model.AddFollow(userID, int64(toUserID))
	if err == nil {
		model.AddFollowCount(userID, 1)
		err := model.AddFollower(int64(toUserID), userID)
		if err == nil {
			model.AddFollowerCount(int64(toUserID), 1)
		}
	}
}

func CancelFollow(userID int64, c *gin.Context) {
	toUserID, _ := strconv.Atoi(c.Query("to_user_id"))
	if userID == int64(toUserID) {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "无法取关自己"})
		return
	}
	err := model.DeleteFollow(userID, int64(toUserID))
	if err == nil {
		model.AddFollowCount(userID, -1)
		err := model.DeleteFollower(int64(toUserID), userID)
		if err == nil {
			model.AddFollowerCount(int64(toUserID), -1)
		}
	}
}

func FollowList(userID int64) []dao.User {
	follows, _ := model.GetFollowList(userID)
	// 当前用户
	//user, _ := model.GetUserByID(userID)
	//follows = append(follows, user)

	for i, _ := range follows {
		follows[i].IsFollow = true
	}
	return follows
}

func FollowerList(userID int64) []dao.User {
	followers, _ := model.GetFollowerList(userID)
	// 将粉丝中，我所关注的人设为IsFollow
	isFollow := make(map[int64]bool)
	follows, _ := model.GetFollowList(userID)
	for _, follow := range follows {
		isFollow[follow.ID] = true
	}
	// 当前用户
	//user, _ := model.GetUserByID(userID)
	//followers = append(followers, user)

	for i, follower := range followers {
		if isFollow[follower.ID] {
			followers[i].IsFollow = true
		}
	}
	return followers
}
