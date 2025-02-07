package model

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Nickname string `gorm:"size:255" json:"nickname"`
	Bio      string `gorm:"type:text; default:no bio yet" json:"bio"`

	Username string `gorm:"size:255;unique;not null" json:"username"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	Role     string `gorm:"size:255;not null;default:user" json:"role"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
}
