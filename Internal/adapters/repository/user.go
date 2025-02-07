package repository

import (
	"errors"
	"strings"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/database"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/pkg/text"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User model.User

var (
	ErrUnableCreateUser = errors.New("unable to create user")
	ErrUserExists = errors.New("user already exists")
)

func (p User) Create() error {
	var (
		userModel model.User
		db = database.GetDatabase()

		name = strings.TrimSpace(p.Name)
		email = strings.TrimSpace(p.Email)
		cleanEmail = text.CleanEmail(email)
		password = strings.TrimSpace(p.Password)
	)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = db.Where("email = ?", cleanEmail).First(&userModel).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		userModel = model.User{
			Name: name,
			Email: cleanEmail,
			Password: string(hashPassword),
		}
	
		if err = db.Create(userModel).Error; err != nil {
			return ErrUnableCreateUser
		}

		return nil
	}
	if err == nil {
		return ErrUserExists
	}

	return err
}