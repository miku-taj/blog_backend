package model

type PostLike struct {
	UserID int `gorm:"primaryKey"`
	PostID int `gorm:"primaryKey"`
}
