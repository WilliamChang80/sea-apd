package postgres

import (
	"github.com/williamchang80/sea-apd/domain"

	"github.com/jinzhu/gorm"
)

// UserRepository is contract for repo User
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}
	return &UserRepository{db: db}
}

// CreateUser ...
func (u *UserRepository) CreateUser(user domain.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByEmail ...
func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
