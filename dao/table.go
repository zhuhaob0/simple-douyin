package dao

import (
	"time"
)

type UsrToken struct {
	Token   string
	EndTime time.Time
}

// TokenValidTime token存活时间
var TokenValidTime = 7 * 24 * time.Hour

// TokenEndTime
// 用户名映射{token,过期时间}
var TokenEndTime = map[string]UsrToken{
	"admin": {
		Token:   "",
		EndTime: time.Now().Add(TokenValidTime),
	},
}

// UsersLoginInfo
// token 映射 User
var UsersLoginInfo = map[string]User{
	//"zhangleidouyin": {
	//	ID:            1,
	//	Name:          "zhanglei",
	//	FollowCount:   10,
	//	FollowerCount: 5,
	//	IsFollow:      true,
	//},
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
