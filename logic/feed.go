package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Feed(token string, c *gin.Context) {
	isValid := TokenIsValid(token)
	if isValid == false { // token 失效
		/*
			此处应该要让用户强制下线，但是没有找到相关接口，所以只是返回一个json
			这个json似乎也没有作用，并不会再客户端提示重新登录的信息
		*/
		log.Println("token失效")
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "登录过期，请重新登录"})
	}
	/*
		目前视频数量较少，所以返回给用户所有视频
		视频增多后，可以对不同种类视频分类，根据用户喜好返回
	*/
	recommendVideo, _ := model.GetAllVideos()

	// 返回的视频中有点赞过的，则设 IsFavorite=true
	user := dao.UsersLoginInfo[token]
	likeVideos, _ := model.GetFavoriteVideosByID(user.ID)
	isLike := make(map[int64]bool)
	for _, video := range likeVideos {
		isLike[video.ID] = true
	}
	// 返回视频作者有关注的，则设IsFollow=true
	isFollow := make(map[int64]bool)
	followUser, _ := model.GetFollowList(user.ID)
	for _, user := range followUser {
		isFollow[user.ID] = true
	}

	for i, video := range recommendVideo {
		if isLike[video.ID] {
			recommendVideo[i].IsFavorite = true
		}
		if isFollow[video.AuthorID] || video.AuthorID == user.ID {
			recommendVideo[i].Author.IsFollow = true
		}
	}
	c.JSON(http.StatusOK, dao.FeedResponse{
		Response:  dao.Response{StatusCode: 0},
		VideoList: recommendVideo,
		NextTime:  time.Now().Unix(),
	})
}
