package user

import (
	"errors"

	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/user"
)

// MockRepository ...
type MockRepository struct {
	ctrl *gomock.Controller
}

// NewMockRepository ...
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

// CreateUser ...
func (m MockRepository) CreateUser(user domain.User) error {
	var u = domain.User{}
	if user == u {
		return errors.New("Cannot create user")
	}
	return nil
}

// GetUserByEmail ...
func (m MockRepository) GetUserByEmail(email string) (*domain.User, error) {
	if email != "" {
		return &domain.User{
			Username: "username",
			Email:    "email",
			Password: "password",
			Role:     "0",
		}, nil
	}
	return nil, errors.New("Cannot get user by email")
}
