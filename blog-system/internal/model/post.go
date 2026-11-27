package model

import (
	"errors"
	"time"
)

type Post struct {
	PostID      int `gorm:"column:post_id;primaryKey"`
	UserID      int `gorm:"column:author_id"`
	Title       string
	Contents    string
	LikesNumber int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
	IsDraft     bool
}

func (p *Post) ValidatePost() error {
	if len(p.Title) < 1 {
		return errors.New("post title must be at least 1 character long")
	}
	if len(p.Contents) < 1 {
		return errors.New("post contents must be at least 1 character long")
	}
	return nil
}
