package models

type Article struct {
	ID          uint
	Title       string
	Description string
	Image       string
	UserID      uint `gorm:"column:userId"`
	User        UserArticle
}
