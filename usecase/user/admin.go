package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
)

// AdminUsecase ...
type AdminUsecase struct {
	ur domain.UserRepository
}

// ConvertToDomain ...
func ConvertToDomain(a request.Admin) domain.User {
	return domain.User{
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
		Role:     "0",
	}
}

// NewAdminUseCase ...
func NewAdminUseCase(a domain.UserRepository) domain.AdminUsecase {
	return &AdminUsecase{
		ur: a,
	}
}

// RegisterAdmin ...
func (s *AdminUsecase) RegisterAdmin(admin request.Admin) error {
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
