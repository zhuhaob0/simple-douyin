package model

import (
	"github.com/MiniDouyin/dao"
	"log"
)

func GetUserByName(username string) (dao.User, error) {
	var user dao.User
	err := dao.Mysql.Debug().Where("name=?", username).Take(&user).Error
	if err != nil {
		log.Println("GetUserByName: ", err)
	}
	return user, err
}

func IsUserExist(username string) bool {
	var user dao.User
	err := dao.Mysql.Debug().Where("name=?", username).Take(&user).Error
	if err != nil {
		log.Println("IsUserExist: ", err)
	}
	return err == nil
}

func GetUserIDByName(username string) int64 {
	user, _ := GetUserByName(username)
	return user.ID
}

func GetAllUser() ([]dao.User, error) {
	var userList []dao.User
	err := dao.Mysql.Debug().Find(&userList).Error //select * from `userinfo`
	if err != nil {
		log.Println("GetAllUser: ", err)
	}
	return userList, err
}

func GetUserByID(id int64) (dao.User, error) {
	var user dao.User
	err := dao.Mysql.Debug().Where("id=?", id).Take(&user).Error
	if err != nil {
		log.Println("GetUserByID: ", err)
	}
	return user, err
}

func GetVideosByAuthorID(id int64) ([]dao.Video, error) {
	var videos []dao.Video
	err := dao.Mysql.Debug().Where("authorID=?", id).Find(&videos).Error
	if err != nil {
		log.Println("GetAllUserInfo: ", err)
	}
	return videos, err
}

func GetAllVideos() ([]dao.Video, error) {
	var allVideos []dao.Video
	// 预加载Author，否则返回的Author为空
	// https://gorm.io/zh_CN/docs/preload.html
	err := dao.Mysql.Debug().Preload("Author").Find(&allVideos).Error
	if err != nil {
		log.Printf("GetAllVideos: %s", err.Error())
	}
	return allVideos, err
}

func GetFavoriteVideosByID(userID int64) ([]dao.Video, error) {
	user := dao.User{ID: userID}
	var favorites []dao.Video
	err := dao.Mysql.Debug().Model(&user).Preload("Author").Association("LikeVideos").Find(&favorites)
	return favorites, err
}

func GetCommentListByVideoID(videoID int64) ([]dao.Comment, error) {
	var comments []dao.Comment
	err := dao.Mysql.Debug().Where("videoID=?", videoID).
		Preload("User").Order("created_at desc").Find(&comments).Error
	return comments, err
}

func GetFollowList(userID int64) ([]dao.User, error) {
	user := dao.User{ID: userID}
	var follows []dao.User
	err := dao.Mysql.Debug().Model(&user).Association("FollowList").Find(&follows)
	return follows, err
}

func GetFollowerList(userID int64) ([]dao.User, error) {
	user := dao.User{ID: userID}
	var followers []dao.User
	err := dao.Mysql.Debug().Model(&user).Association("FollowerList").Find(&followers)
	return followers, err
}
