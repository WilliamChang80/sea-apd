package user

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
	"github.com/williamchang80/sea-apd/dto/response/base"

	"net/http"
)

// AdminController ...
type AdminController struct {
	usecase user.AdminUsecase
}

// NewAdminController ...
func NewAdminController(e *echo.Echo, a user.AdminUsecase) user.AdminController {
	c := &AdminController{
		usecase: a,
	}
	e.POST("api/user/admin", c.RegisterAdmin)

	return c
}

// RegisterAdmin ...
func (a *AdminController) RegisterAdmin(c echo.Context) error {
	var adminRequest admin.Admin
	c.Bind(&adminRequest)

	if err := a.usecase.RegisterAdmin(adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}
