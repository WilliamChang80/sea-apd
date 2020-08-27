package repository

import (
	"github.com/williamchang80/sea-apd/domain/user"

	"github.com/jinzhu/gorm"
)

// UserRepository is contract for repo User
type UserRepository struct {
	db *gorm.DB
}

// User ...
var User user.User

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&user.User{})
	return UserRepository{db: db}
}

// CreateUser ...
func (u *UserRepository) CreateUser(user *user.User) (user.User, error) {
	u.db.Create(&user)
	return User, nil
}
