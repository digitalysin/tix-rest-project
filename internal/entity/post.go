package entity

import "time"

type (
	Post struct {
		Id        int64     `gorm:"primaryKey"`
		Title     string    `gorm:"column:title"`
		Content   string    `gorm:"column:content"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}
)

func (impl *Post) TableName() string {
	return "posts"
}
