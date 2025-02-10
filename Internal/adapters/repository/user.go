package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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
	ErrNoUser = errors.New("user dose not exists")

	ErrEncodeJson = errors.New("unable to encode data to JSON")
	ErrDecodeJson = errors.New("unable to decode JSON to data")
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

	err = db.Where("email = ?", p.Data.Email).First(&userModel).Error
	
	if err == gorm.ErrRecordNotFound {
		userModel = p.Data
		if err = db.Create(&userModel).Error; err != nil {
			return ErrUnableCreateUser
		}

		if err := setUserCache(ctx, userModel); err != nil {
			return err
		}

		p.AddData(userModel)

		return nil
	}
	if err == nil {
		return ErrUserExists
	}

	return err
}

func (p User) GetbyUsername() error {
	var (
		ctx = context.Background()
		db = database.GetDatabase()

		username = strings.TrimSpace(p.Data.Username)
	)

	if err := getUserCachebyUsername(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			return ErrNoUser
		}
		if err = db.Where("username = ?", username).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				setUserCacheNoData(ctx, p.Data)
			}
			return ErrUnableGetUser
		}
	}

	return nil
}

// =========== Lib ============

func setUserCache(ctx context.Context, user model.User) error {
	var (
		keyname = fmt.Sprintf("user:%s", user.Username)
		keyid = fmt.Sprintf("userid:%d", user.ID)
	)

	data, err := json.Marshal(user)
	if err != nil {
		return ErrEncodeJson
	}

	if err = cache.Set(ctx, keyname, string(data)); err != nil {
		return err
	}
	if err = cache.Set(ctx, keyid, string(data)); err != nil {
		return err
	}

	return nil
}

func setUserCacheNoData(ctx context.Context, user model.User) error {
	var (
		keyname = fmt.Sprintf("user:%s", user.Username)
		keyid = fmt.Sprintf("userid:%d", user.ID)
	)

	if err := cache.SetNoData(ctx, keyname); err != nil {
		return err
	}
	if err := cache.SetNoData(ctx, keyid); err != nil {
		return err
	}

	return nil
}

func getUserCachebyUsername(ctx context.Context, user *model.User) error {
	var (
		key = fmt.Sprintf("user:%s", user.Username)
	)
	data, err := cache.Get(ctx, key)
	if err != nil {
		return err
	}

	var cachedUser model.User
	if err := json.Unmarshal([]byte(data), &cachedUser); err != nil {
		return ErrDecodeJson
	}

	*user = cachedUser

	return nil
}