package dao

// 一个视频有一个作者 one2one-belongs to
// 一个视频有多个点赞的人，一个人点赞多个视频 many2many
// 一个视频有多条评论 one2many

type Video struct {
	ID            int64     `json:"id,omitempty" gorm:"column:id;unique;primaryKey;autoIncrement"`
	AuthorID      int64     `json:"authorID,omitempty" gorm:"column:authorID"`
	Author        User      `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	PlayUrl       string    `json:"play_url,omitempty" gorm:"column:playUrl"`
	CoverUrl      string    `json:"cover_url,omitempty" gorm:"column:coverUrl"`
	FavoriteCount int64     `json:"favorite_count,omitempty" gorm:"column:favoriteCount"`
	CommentCount  int64     `json:"comment_count,omitempty" gorm:"column:commentCount"`
	IsFavorite    bool      `json:"is_favorite,omitempty" gorm:"column:isFavorite"`
	Title         string    `json:"title,omitempty" gorm:"column:title"`
	Likers        []User    `gorm:"many2many:Favorite"`
	CommentList   []Comment `gorm:"foreignKey:videoID"`
}

func (v Video) TableName() string {
	return "video"
}
