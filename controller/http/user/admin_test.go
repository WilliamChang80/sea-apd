package user

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"

	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
	user_mock_repository "github.com/williamchang80/sea-apd/mocks/repository/user"
	user_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/user"
	"github.com/williamchang80/sea-apd/usecase/user"
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
	ctx := echo.New()
	repo := user_mock_repository.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}

	tests := []struct {
		name     string
		args     args
		want     domain.AdminController
		initMock func() domain.AdminUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &AdminController{
				usecase: user.NewAdminUseCase(repo),
			},
			initMock: func() domain.AdminUsecase {
				c := user_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()

			if got := NewAdminController(tt.args.ctx, mock); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewAdminController() = %v, want %v", got, tt.want)
			}
		})
	}
}
