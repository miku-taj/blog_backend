package model

type Subscription struct {
	UserID   int `gorm:"primaryKey"`
	AuthorID int `gorm:"primaryKey"`
}
