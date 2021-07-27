package bean

import (
	"time"
)

type News struct {
	ID        uint      `json:"id "gorm:"primaryKey"`
	Title     string    `form:"title" json:"title" binding:"required"`
	Content   string    `form:"contetn" json:"content" binding:"required"`
	Url       string    `json:"url"`
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli"`
}
