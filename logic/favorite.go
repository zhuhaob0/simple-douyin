package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(token string, c *gin.Context) {
	// actionType := c.PostForm("action_type")
	// 此处是POST请求，action_type不应该附在URL上，可能前端出错
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	user := dao.UsersLoginInfo[token]

	if actionType == 1 { // 点赞
		err := model.AddLikeVideo(user.ID, int64(videoID))
		if err == nil {
			model.AddFavoriteCount(int64(videoID), 1)
		}
	} else {
		err := model.DeleteLikeVideo(user.ID, int64(videoID))
		if err == nil {
			model.AddFavoriteCount(int64(videoID), -1)
		}
	}
}

func FavoriteList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	favorites, _ := model.GetFavoriteVideosByID(int64(userID))
	c.JSON(http.StatusOK, dao.VideoListResponse{
		Response: dao.Response{
			StatusCode: 0,
		},
		VideoList: favorites,
	})
}
