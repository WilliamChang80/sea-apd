package user

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/williamchang80/sea-apd/dto/request/admin"
)

var (
	mockData = admin.Admin{
		Username: "admin",
		Email:    "admin@admin.com",
		Password: "p4ssw0rd",
	}
	mockID = "1"
)

func TestNewAdminController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// ctx := echo.New()
}
