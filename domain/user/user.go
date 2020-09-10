package user

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/common/security"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/admin"
	"github.com/williamchang80/sea-apd/dto/request/auth"
)

// User ...
type User struct {
	domain.Base
	Name     string `gorm:"size:50;not null;" json:"name"`
	Email    string `gorm:"unique;size:100;not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Role     string `gorm:"size:1;not null;" json:"role"`
}

// UserRepository ...
type UserRepository interface {
	CreateUser(User) error
	GetUserByEmail(string) (*User, error)
}

// AdminUsecase ...
type AdminUsecase interface {
	RegisterAdmin(admin.Admin) error
}

type UserUsecase interface {
	CreateUser(request auth.RegisterUserRequest) error
}

// AdminController ...
type AdminController interface {
	RegisterAdmin(echo.Context) error
}

type UserController interface {
	CreateUser(echo.Context) error
}

// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)

	return nil
}
