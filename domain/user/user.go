package user

import (
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/common/security"
	"github.com/williamchang80/sea-apd/domain/base"
)

// User ...
type User struct {
	*base.Base
	Name     string `gorm:"size:50;not null;" json:"name"`
	Email    string `gorm:"unique;size:100;not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Role     string `gorm:"size:1;not null;" json:"role"`
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

// Prepare for sanitize input
func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate for input validation
func (u *User) Validate(action string) string {
	var err error

	switch strings.ToLower(action) {
	default:
		if u.Name == "" {
			return "name is required"
		}
		if u.Email == "" {
			return "email is required"
		}
		if err = checkmail.ValidateFormat(u.Email); err != nil {
			return "please provide a valid email"
		}
		if u.Password == "" {
			return "password is required"
		}
		if len(u.Password) < 6 {
			return "password should be at least 6 characters"
		}
	}

	return "something wrong"
}
