package routes

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	use_case "github.com/williamchang80/sea-apd/usecase/user"
)

// NewAdminRoutes ...
func NewAdminRoutes(e *echo.Echo) {
	db := db.Postgres()
	repo := postgres.NewUserRepository(db)
	usecase := use_case.NewAdminUseCase(repo)
	user.NewAdminController(e, usecase)
}
