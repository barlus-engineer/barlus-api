package services

import (
	"fmt"
	"strings"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/pkg/text"
)

type User struct {
	Data model.User
}

type UserForm struct {
	Name     string
	Nickname string
	Bio      string
	Username string
	Email    string
	Password string
}

func (p *User) AddData(data model.User) {
	n_name := strings.TrimSpace(data.Name)
	n_nickname := strings.TrimSpace(data.Nickname)
	n_username := text.CleanUsername(strings.TrimSpace(data.Username))
	n_email := text.CleanEmail(strings.TrimSpace(data.Email))
	n_password := text.HashPassword(strings.TrimSpace(data.Password))

	p.Data = model.User{
		Name:     n_name,
		Nickname: n_nickname,
		Bio:      data.Bio,
		Username: n_username,
		Email:    n_email,
		Password: n_password,
	}
}

func (p User) Register() error {
	var (
		userRepo repository.User
		err error
	)

	fmt.Println(p.Data)
	userRepo.AddData(p.Data)

	if err = userRepo.Create(); err != nil {
		return err
	}

	return nil
}

// === lib ===

