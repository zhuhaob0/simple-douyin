package logic

import (
	"fmt"
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
)

func Comment(user *dao.User, c *gin.Context) {
	text := c.Query("comment_text")
	if utf8.RuneCountInString(text) > 64 {
		c.JSON(http.StatusOK, dao.CommentActionResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "评论字数超过64字符"},
		})
		return
	}
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	comment := dao.Comment{
		UserID:     user.ID,
		VideoID:    int64(videoID),
		User:       *user,
		Content:    text,
		CreateDate: fmt.Sprintf("%02d-%02d", time.Now().Month(), time.Now().Day()),
	}
	newID := model.AddComment(&comment)
	//println("newID=", newID)
	if newID != 0 {
		model.AddCommentCount(comment.VideoID, 1)
	}
	comment.ID = newID
	c.JSON(http.StatusOK, dao.CommentActionResponse{
		Response: dao.Response{StatusCode: 0, StatusMsg: "评论成功"},
		Comment:  comment,
	})
}

func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Query("comment_id"))
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	err := model.DeleteComment(int64(commentID))
	if err == nil {
		model.AddCommentCount(int64(videoID), -1)
	}
	c.JSON(http.StatusOK, dao.CommentActionResponse{
		Response: dao.Response{StatusCode: 0, StatusMsg: "删除成功"},
	})
}

func CommentList(c *gin.Context) []dao.Comment {
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	comments, _ := model.GetCommentListByVideoID(int64(videoID))
	return comments
}
