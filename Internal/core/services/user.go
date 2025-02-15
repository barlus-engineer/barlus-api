package services

import (
	"errors"
	"strings"

	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"github.com/barlus-engineer/barlus-api/Internal/ports"
	"github.com/barlus-engineer/barlus-api/pkg/text"
)

var (
	ErrUsernameExists = errors.New("this username is already exists")
	ErrEmailExists = errors.New("this email is already exists")
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

func (p *UserService) AddData(data model.User) {
	p.Data = data
}

func (p UserService) Register(data dto.UserRegisterRequest) error {
	var (
		err error
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

	if err = p.Repo.AddData(p.Data).Create().Error(); err != nil {
		return err
	}

	return nil
}

func (p UserService) UsernameAvail(data dto.UserUsernameAvailRequest) error {
	var (
		err error
	)

	n_username := text.CleanUsername(strings.TrimSpace(data.Username))

	p.Data = model.User{
		Username: n_username,
	}

	if err = p.Repo.AddData(p.Data).GetbyUsername().ReturnData(&p.Data).Error(); err != nil {
		return nil
	}

	return ErrUsernameExists
}

func (p *UserService) EmailAvail(data dto.UserEmailAvailRequest) error {
	var (
		err error
	)

	n_email := text.CleanEmail(strings.TrimSpace(data.Email))

	p.Data = model.User{
		Email: n_email,
	}

	if err = p.Repo.AddData(p.Data).GetbyEmail().ReturnData(&p.Data).Error(); err != nil {
		return nil
	}

	return ErrEmailExists
}