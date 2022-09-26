package model

import (
	"github.com/MiniDouyin/dao"
	"log"
)

// AddIntoUserDB
// 添加User进数据库user表
func AddIntoUserDB(user dao.User) {
	err := dao.Mysql.Debug().Create(&user).Error
	if err != nil {
		log.Println("AddIntoUserDB: ", err)
		return
	}
}

// AddIntoVideo
// 将publish的视频保存进数据库
func AddIntoVideo(video dao.Video) {
	err := dao.Mysql.Debug().Create(&video).Error
	if err != nil {
		log.Println("AddIntoVideo: ", err)
		return
	}
}

// AddLikeVideo 点赞操作
func AddLikeVideo(userID, videoID int64) error {
	user := dao.User{ID: userID}
	video := dao.Video{ID: videoID, IsFavorite: true}
	err := dao.Mysql.Debug().Model(&user).Association("LikeVideos").Append(&video)
	return err
}

// AddComment 评论操作
func AddComment(comment *dao.Comment) int64 {
	video := dao.Video{ID: comment.VideoID}
	err := dao.Mysql.Debug().Model(&video).Association("CommentList").Append(comment)
	if err != nil {
		return 0
	}
	// 并发高时，要查询的最新的评论可能不是此函数内创建的那个，就会出错
	var newestComment dao.Comment
	err = dao.Mysql.Debug().Order("created_at desc").First(&newestComment).Error
	if err != nil {
		return 0
	}
	return newestComment.ID
}

// AddFollow 添加关注
func AddFollow(userID, followID int64) error {
	user := dao.User{ID: userID}
	follow := dao.User{ID: followID}
	err := dao.Mysql.Debug().Model(&user).Association("FollowList").Append(&follow)
	return err
}

// AddFollower 添加粉丝
func AddFollower(userID, followerID int64) error {
	user := dao.User{ID: userID}
	follower := dao.User{ID: followerID}
	err := dao.Mysql.Debug().Model(&user).Association("FollowerList").Append(&follower)
	return err
}
