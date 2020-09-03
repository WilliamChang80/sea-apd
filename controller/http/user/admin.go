package user

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mid "github.com/williamchang80/sea-apd/controller/middleware"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
	"github.com/williamchang80/sea-apd/dto/response/base"

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
	e.Use(middleware.BasicAuth(mid.BasicAuthAdmin))

	return c
}

// RegisterAdmin ...
func (a *AdminController) RegisterAdmin(c echo.Context) error {
	var adminRequest request.Admin
	c.Bind(&adminRequest)

	if err := a.usecase.RegisterAdmin(adminRequest); err != nil {
		if err.Error() == "duplicate" {
			return c.JSON(http.StatusBadRequest, &base.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Email has been taken",
			})
		}
		return c.JSON(http.StatusInternalServerError, &base.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Error",
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Admin created successfully",
	})
}
