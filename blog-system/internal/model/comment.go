package model

import (
	"database/sql"
	"time"
)

type Comment struct {
	CommentID   int `gorm:"column:comment_id;primaryKey"`
	UserID      int
	PostID      int
	ParentID    sql.NullInt64
	Content     string
	LikesNumber int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}
