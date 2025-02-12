package ports

import (
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"gorm.io/gorm"
)

type UserRepo interface {
	AddDatabase(database *gorm.DB)
	AddData(data model.User)
	Create() error
	GetbyID() error
	GetbyUsername() error
	GetbyEmail() error
}

type UserService interface {
	Register(data dto.UserRegisterForm) error
}