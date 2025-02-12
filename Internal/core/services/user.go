package services

import (
	"strings"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/database"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"github.com/barlus-engineer/barlus-api/Internal/ports"
	"github.com/barlus-engineer/barlus-api/pkg/text"
)

type UserService struct {
	Repo	ports.UserRepo
	Data    model.User
}

func  NewUserService(repo ports.UserRepo) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (p UserService) Register(data dto.UserRegisterForm) error {
	var (
		Database = database.GetDatabase()
		err      error
	)

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

	p.Repo.AddDatabase(Database)
	p.Repo.AddData(p.Data)

	if err = p.Repo.Create(); err != nil {
		return err
	}

	return nil
}

// === lib ===
