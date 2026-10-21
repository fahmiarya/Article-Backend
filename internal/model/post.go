package model

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
}

func (Post) TableName() string {
	return "posts"
}
