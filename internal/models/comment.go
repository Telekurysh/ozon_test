package models

import "time"

type Comment struct {
	ID        int `gorm:"primaryKey"`
	PostID    int `gorm:"index"`
	ParentID  *int
	Content   string `gorm:"type:text;size:2000"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Children  []Comment `gorm:"foreignKey:ParentID"`
}
