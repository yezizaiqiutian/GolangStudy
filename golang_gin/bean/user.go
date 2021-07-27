package bean

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserName    string         `form:"user_name" json:"user_name" binding:"required"`
	PassWord    string         `form:"pass_word" json:"pass_word" binding:"required"`
	Email       *string        `json:"email"`
	Age         uint8          `json:"age"`
	Birthday    *time.Time     `json:"birthday"`
	Number      sql.NullString `json:"number"`
	ActivatedAt sql.NullTime   `json:"activated_at"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoUpdateTime:milli"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime:milli"`
}
