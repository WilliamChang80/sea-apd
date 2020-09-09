package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
)

// AdminUsecase ...
type AdminUsecase struct {
	ur user.UserRepository
}

// ConvertToDomain ...
func ConvertToDomain(a admin.Admin) user.User {
	return user.User{
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
		Role:     "0",
	}
}

// NewAdminUseCase ...
func NewAdminUseCase(a user.UserRepository) user.AdminUsecase {
	return &AdminUsecase{
		ur: a,
	}
}

// RegisterAdmin ...
func (s *AdminUsecase) RegisterAdmin(admin admin.Admin) error {
	u, err := s.ur.GetUserByEmail(admin.Email)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if u != nil {
		return errors.New("duplicate")
	}

	a := ConvertToDomain(admin)
	err = s.ur.CreateUser(a)
	if err != nil {
		return err
	}
	return nil
}
