package model

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Nickname string `gorm:"size:255"`
	Bio      string `gorm:"type:text; default:no bio yet"`

	Username  string	`gorm:"size:255;unique;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"size:255;not null;default:user"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
