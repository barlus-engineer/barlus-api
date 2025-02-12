package ports

import (
	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"gorm.io/gorm"
)

type UserRepo interface {
	AddData(data model.User) repository.UserRepo
	AddDatabase(database *gorm.DB) repository.UserRepo
	Create() repository.UserRepo
	Error() error
	GetbyEmail() repository.UserRepo
	GetbyID() repository.UserRepo
	GetbyUsername() repository.UserRepo
	ReturnData(userModel *model.User) repository.UserRepo
}

type UserService interface {
	Register(data dto.UserRegisterRequest) error
	UsernameAvail(data dto.UserUsernameAvailRequest) error
}