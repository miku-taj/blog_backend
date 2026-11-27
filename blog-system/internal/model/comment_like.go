package model

type CommentLike struct {
	UserID    int `gorm:"primaryKey"`
	CommentID int `gorm:"primaryKey"`
}
