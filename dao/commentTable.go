package dao

import "gorm.io/gorm"

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// 一个视频有多条评论 one2many
// 一条评论有一个作者 one2one-belongs to

type Comment struct {
	gorm.Model
	ID         int64  `json:"id,omitempty" gorm:"column:id;unique;primaryKey"`
	UserID     int64  `json:"userid,omitempty" gorm:"column:userID"`
	VideoID    int64  `json:"video_id,omitempty" gorm:"column:videoID"`
	User       User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Content    string `json:"content,omitempty" gorm:"column:content"`
	CreateDate string `json:"create_date,omitempty" gorm:"column:createDate;not null;"`
}

func (comment Comment) TableName() string {
	return "comment"

}
