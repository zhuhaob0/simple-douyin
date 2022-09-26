package controller

import (
	"github.com/MiniDouyin/dao"
)

var DemoVideos = []dao.Video{
	{
		ID:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.0.4:8080/static/demo/程咬金.mp4",
		CoverUrl:      "http://192.168.0.4:8080/static/demo/程咬金.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		ID:            2,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.0.4:8080/static/demo/明世隐.mp4",
		CoverUrl:      "http://192.168.0.4:8080/static/demo/明世隐.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    true,
	},
}

//var DemoComments = []dao.Comment{
//	{
//		ID:         1,
//		User:       DemoUser,
//		Content:    "Test Comment",
//		CreateDate: "05-01",
//	},
//}

var DemoUser = dao.User{
	ID:            0,
	Name:          "admin",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
