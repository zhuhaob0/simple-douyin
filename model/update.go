package model

import (
	"github.com/MiniDouyin/dao"
	"gorm.io/gorm"
	"log"
)

// SetOnline 根据id设置用户的在线状态，set online = status
func SetOnline(id int64, status bool) error {
	user := dao.User{}
	err := dao.Mysql.Debug().Where("id=?", id).Take(&user).Error
	if err != nil {
		log.Println("SetOnline: ", err)
		return err
	}
	//通过User模型的主键id的值作为where条件，更新online字段值
	dao.Mysql.Model(&user).Debug().Update("online", status)
	return nil
}

// AddFavoriteCount 修改视频点赞数量
func AddFavoriteCount(videoID, addNum int64) {
	video := dao.Video{ID: videoID}
	err := dao.Mysql.Debug().Model(&video).
		Update("favoriteCount", gorm.Expr("favoriteCount+?", addNum)).Error
	if err != nil {
		return
	}
}

// AddCommentCount 修改评论数量
func AddCommentCount(videoID, addNum int64) {
	video := dao.Video{ID: videoID}
	err := dao.Mysql.Debug().Model(&video).
		Update("commentCount", gorm.Expr("commentCount+?", addNum)).Error
	if err != nil {
		return
	}
}

// AddFollowCount 修改关注数量
func AddFollowCount(userID, addNum int64) {
	user := dao.User{ID: userID}
	err := dao.Mysql.Debug().Model(&user).
		Update("followCount", gorm.Expr("followCount+?", addNum)).Error
	if err != nil {
		return
	}
}

// AddFollowerCount 修改粉丝数量
func AddFollowerCount(userID, addNum int64) {
	user := dao.User{ID: userID}
	err := dao.Mysql.Debug().Model(&user).
		Update("followerCount", gorm.Expr("followerCount+?", addNum)).Error
	if err != nil {
		return
	}
}
