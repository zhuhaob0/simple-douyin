package dao

import "gorm.io/gorm"

// User 用户信息
type User struct {
	ID            int64   `json:"id,omitempty" gorm:"column:id;unique;primaryKey;"`
	Name          string  `json:"name,omitempty" gorm:"column:name;unique;"`
	FollowCount   int64   `json:"follow_count,omitempty" gorm:"column:followCount"`
	FollowerCount int64   `json:"follower_count,omitempty" gorm:"column:followerCount"`
	IsFollow      bool    `json:"is_follow,omitempty" gorm:"column:isFollow"`
	Password      string  `gorm:"column:password"`
	Online        bool    `gorm:"column:online"`
	LikeVideos    []Video `gorm:"many2many:Favorite"`
	FollowList    []*User `gorm:"many2many:follow;"`
	FollowerList  []*User `gorm:"many2many:follower;"`
	gorm.Model
}

func (user User) TableName() string {
	return "user"
}
