package user

import (
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/repository"
)

// AdminUseCase ...
type AdminUseCase interface {
	RegisterAdmin(user user.User) (user.User, error)
}

// AdminUseCaseImpl ...
type AdminUseCaseImpl struct {
	ur repository.UserRepository
}

// NewAdminUseCaseImpl ...
func NewAdminUseCaseImpl(a repository.UserRepository) AdminUseCase {
	return &AdminUseCaseImpl{
		ur: a,
	}
}

// RegisterAdmin ...
func (s AdminUseCaseImpl) RegisterAdmin(user user.User) (user.User, error) {
	return s.ur.CreateUser(&user)
}
