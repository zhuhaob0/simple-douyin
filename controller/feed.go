package controller

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	//token为空说明用户未登录，返回demo数据
	if token == "" {
		c.JSON(http.StatusOK, dao.FeedResponse{
			Response:  dao.Response{StatusCode: 0},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	logic.Feed(token, c)
}
