package user

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"

	"net/http"
)

// AdminController ...
type AdminController struct {
	usecase domain.AdminUsecase
}

// NewAdminController ...
func NewAdminController(e *echo.Echo, a domain.AdminUsecase) domain.AdminController {
	c := &AdminController{
		usecase: a,
	}
	e.POST("/users/register-admin", c.RegisterAdmin)
	return c
}

// RegisterAdmin ...
func (a *AdminController) RegisterAdmin(c echo.Context) error {
	var adminRequest request.Admin
	c.Bind(&adminRequest)

	if err := a.usecase.RegisterAdmin(adminRequest); err != nil {
		if err.Error() == "duplicate" {
			return c.JSON(http.StatusBadRequest, "Email has been taken")
		}
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, "Admin created successfully")
}
