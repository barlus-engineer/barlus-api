package repository

import (
	"context"
	"errors"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/cache"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"gorm.io/gorm"
)

type UserRepo struct {
	Database gorm.DB
	Data     model.User
	err      error
}

var (
	ErrUnableCreateUser = errors.New("unable to create user")
	ErrUnableGetUser    = errors.New("unable to get user")
	ErrUserExists       = errors.New("user already exists")
	ErrUserEmailExists  = errors.New("user email already exists")
	ErrNoUser           = errors.New("user dose not exists")
)

var (
	logUnableCreateUser        = "repo/user error create user '%s', email '%s'"
	logUnableGetUserbyID       = "repo/user (get by id[%d]): %s"
	logUnableGetUserbyUsername = "repo/user (get by username[%s]): %s"
	logUnableGetUserbyEmail    = "repo/user (get by email[%s]): %s"
)

func (p *UserRepo) AddDatabase(database *gorm.DB) UserRepo {
	p.Database = *database
	return *p
}

func (p *UserRepo) AddData(data model.User) UserRepo {
	p.Data = data
	return *p
}

func (p UserRepo) Error() error {
	return p.err
}

func (p UserRepo) ReturnData(userModel *model.User) UserRepo {
	*userModel = p.Data
	return p
}

func (p UserRepo) Create() UserRepo {
	var (
		ctx       = context.Background()
		db        = p.Database
		userModel model.User
		err       error
	)

	if err = p.GetbyUsername().Error(); err == nil {
		p.err = ErrUserExists
		return p
	}
	if err = p.GetbyEmail().Error(); err == nil {
		p.err = ErrUserEmailExists
		return p
	}

	if p.err == ErrNoUser {
		userModel = p.Data
		if err = db.Create(&userModel).Error; err != nil {
			logger.Warningf(logUnableCreateUser, userModel.Username, userModel.Email)
			p.err = ErrUnableCreateUser
			return p
		}
		if err := cache.SetUserCache(ctx, userModel); err != nil {
			p.err = err
			return p
		}
		p.AddData(userModel)
		return p
	}
	return p
}

func (p UserRepo) GetbyID() UserRepo {
	var (
		ctx = context.Background()
		db  = p.Database
		ID  = p.Data.ID
	)
	if err := cache.GetUserbyUsername(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			p.err = ErrNoUser
			return p
		}
		if err = db.Where("id = ?", ID).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err = cache.SetUserIDNotfound(ctx, p.Data); err != nil {
					logger.Alert("repo/user (cache set user id notfound): ", err.Error())
				}
				p.err = ErrNoUser
				return p
			}
			logger.Alertf(logUnableGetUserbyID, p.Data.ID, err.Error())
			p.err = ErrUnableGetUser
			return p
		}
	}
	return p
}

func (p UserRepo) GetbyUsername() UserRepo {
	var (
		ctx      = context.Background()
		db       = p.Database
		username = p.Data.Username
	)
	if err := cache.GetUserbyUsername(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			p.err = ErrNoUser
			return p
		}
		if err = db.Where("username = ?", username).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err = cache.SetUserUsernameNotfound(ctx, p.Data); err != nil {
					logger.Alert("repo/user (cache set user username notfound): ", err.Error())
				}
				p.err = ErrNoUser
				return p
			}
			logger.Alert(logUnableGetUserbyUsername, p.Data.Username, err.Error())
			p.err = ErrUnableGetUser
			return p
		}
	}
	return p
}

func (p UserRepo) GetbyEmail() UserRepo {
	var (
		ctx   = context.Background()
		db    = p.Database
		email = p.Data.Email
	)
	if err := cache.GetUserbyEmail(ctx, &p.Data); err != nil {
		if err == cache.ErrNotFound {
			p.err = ErrNoUser
			return p
		}
		if err = db.Where("email = ?", email).First(&p.Data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err = cache.SetUserEmailNotfound(ctx, p.Data); err != nil {
					logger.Alert("repo/user (cache set user email notfound): ", err.Error())
				}
				p.err = ErrNoUser
				return p
			}
			logger.Alert(logUnableGetUserbyEmail, p.Data.Email, err.Error())
			p.err = ErrUnableGetUser
			return p
		}
	}
	return p
}
