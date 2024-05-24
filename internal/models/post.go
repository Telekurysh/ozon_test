package models

import "time"

type Post struct {
	ID               int    `gorm:"primaryKey"`
	Title            string `gorm:"size:255"`
	Content          string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CommentsDisabled bool
	Comments         []Comment `gorm:"foreignKey:PostID"`
}
