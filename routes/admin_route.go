package routes

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/user"
	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	repo "github.com/williamchang80/sea-apd/repository/postgres/user"
	use_case "github.com/williamchang80/sea-apd/usecase/user"
)

type AdminRoute struct {
	controller domain.AdminController
	usecase    domain.AdminUsecase
	Repository domain.UserRepository
}

// NewAdminRoutes ...
func NewAdminRoutes(e *echo.Echo) AdminRoute {
	db := db.Postgres()
	authRoute := NewAuthRoute(e)
	repo := repo.NewUserRepository(db)
	usecase := use_case.NewAdminUseCase(repo, authRoute.usecase)
	controller := user.NewAdminController(e, usecase)
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}

	return AdminRoute{
		controller: controller,
		usecase:    usecase,
		Repository: repo,
	}
}
