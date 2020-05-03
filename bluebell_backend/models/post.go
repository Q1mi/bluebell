package models

import "time"

type Post struct {
	PostID      uint64    `json:"post_id" db:"post_id"`
	Caption     string    `json:"caption" db:"caption"`
	Content     string    `json:"content" db:"content"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
