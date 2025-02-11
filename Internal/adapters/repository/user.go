package repository

import (
	"context"
	"errors"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/cache"
	"github.com/barlus-engineer/barlus-api/Internal/adapters/database"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"gorm.io/gorm"
)

type User  struct {
	Data model.User
}

var (
	ErrUnableCreateUser = errors.New("unable to create user")
	ErrUnableGetUser = errors.New("unable to get user")
	ErrUserExists = errors.New("user already exists")
	ErrUserEmailExists = errors.New("user email already exists")
	ErrNoUser = errors.New("user dose not exists")
)

func (p *User) AddData(data model.User) {
	p.Data = data
}

func (p User) Create() error {
	var (
		ctx = context.Background()
		db = database.GetDatabase()
		userModel model.User
		err error
	)

	if err = p.GetbyUsername(); err == nil {
		return ErrUserExists
	}
	if err = p.GetbyEmail(); err == nil {
		return ErrUserEmailExists
	}

	if err == ErrNoUser {
		userModel = p.Data
		if err = db.Create(&userModel).Error; err != nil {
			return ErrUnableCreateUser
		}
		if err := cache.SetUserCache(ctx, userModel); err != nil {
			return err
		}
		p.AddData(userModel)
		return nil
	}

	return err
}

func (p User) GetbyID() error {
	var (
		ctx = context.Background()
		db = database.GetDatabase()
		ID = p.Data.ID
	)
	if err := cache.GetUserbyUsername(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			return ErrNoUser
		}
		if err = db.Where("id = ?", ID).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				cache.SetUserIDNotfound(ctx, p.Data)
				return ErrNoUser
			}
			return ErrUnableGetUser
		}
	}
	return nil
}

func (p User) GetbyUsername() error {
	var (
		ctx = context.Background()
		db = database.GetDatabase()
		username = p.Data.Username
	)
	if err := cache.GetUserbyID(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			return ErrNoUser
		}
		if err = db.Where("username = ?", username).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				cache.SetUserUsernameNotfound(ctx, p.Data)
				return ErrNoUser
			}
			return ErrUnableGetUser
		}
	}
	return nil
}

func (p User) GetbyEmail() error {
	var (
		ctx = context.Background()
		db = database.GetDatabase()
		email = p.Data.Email
	)
	if err := cache.GetUserbyEmail(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			return ErrNoUser
		}
		if err = db.Where("email = ?", email).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				cache.SetUserEmailNotfound(ctx, p.Data)
				return ErrNoUser
			}
			return ErrUnableGetUser
		}
	}
	return nil
}