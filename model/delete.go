package model

import (
	"github.com/MiniDouyin/dao"
)

// DeleteLikeVideo 取消点赞
func DeleteLikeVideo(userID, videoID int64) error {
	user := dao.User{ID: userID}
	video := dao.Video{ID: videoID}
	err := dao.Mysql.Debug().Model(&user).Association("LikeVideos").Delete(&video)
	return err
}

// DeleteComment 删除评论
func DeleteComment(commentID int64) error {
	comment := dao.Comment{ID: commentID}
	// 只会删除引用，将videoID置为 null
	//err := dao.Mysql.Debug().Model(&dao.Video{ID: videoID}).Association("CommentList").Delete(&comment)
	//if err != nil {
	//	return
	//}
	err := dao.Mysql.Unscoped().Debug().Delete(&comment).Error
	return err
}

// DeleteFollow 取消关注-删除关注的人
func DeleteFollow(userID, followID int64) error {
	user := dao.User{ID: userID}
	follow := dao.User{ID: followID}
	err := dao.Mysql.Debug().Model(&user).Association("FollowList").Delete(&follow)
	return err
}

// DeleteFollower 取消关注-删除粉丝
func DeleteFollower(userID, followerID int64) error {
	user := dao.User{ID: userID}
	follower := dao.User{ID: followerID}
	err := dao.Mysql.Debug().Model(&user).Association("FollowerList").Delete(&follower)
	return err
}
