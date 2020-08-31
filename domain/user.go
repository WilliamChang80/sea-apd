package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/common/security"
	"github.com/williamchang80/sea-apd/dto/request"
)

// User ...
type User struct {
	Base
	Name     string `gorm:"size:50;not null;" json:"name"`
	Email    string `gorm:"unique;size:100;not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Role     string `gorm:"size:1;not null;" json:"role"`
}

// UserRepository ...
type UserRepository interface {
	CreateUser(User) error
}

// AdminUsecase ...
type AdminUsecase interface {
	RegisterAdmin(request.Admin) error
}

// AdminController ...
type AdminController interface {
	RegisterAdmin(echo.Context) error
}

// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	scope.SetColumn("ID", id)

	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)

	return nil
}
