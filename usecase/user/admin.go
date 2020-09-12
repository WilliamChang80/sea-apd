package user

import (
	"errors"
	"os"

	"github.com/williamchang80/sea-apd/dto/request/auth"

	"github.com/williamchang80/sea-apd/common/constants/user_role"
	auth_domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
)

// AdminUsecase ...
type AdminUsecase struct {
	ur      user.UserRepository
	usecase auth_domain.AuthUsecase
}

// NewAdminUseCase ...
func NewAdminUseCase(a user.UserRepository, usecase auth_domain.AuthUsecase) user.AdminUsecase {
	return &AdminUsecase{
		ur:      a,
		usecase: usecase,
	}
}

// RegisterAdmin ...
func (s *AdminUsecase) RegisterAdmin(request admin.Admin) error {
	authRequest := auth.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
	if _, err := s.usecase.Login(authRequest); err != nil {
		return err
	}

	u, err := s.ur.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("credential not match")
	}

	isValid := validateToken(request.Token)
	if isValid == false {
		return errors.New("Token not valid")
	}

	if err := s.ur.UpdateUserRole(user_role.ToString(user_role.ADMIN), u.ID); err != nil {
		return err
	}

	return nil
}

func validateToken(token string) bool {
	if token == os.Getenv("ADMIN_TOKEN") {
		return true
	}

	return false
}
