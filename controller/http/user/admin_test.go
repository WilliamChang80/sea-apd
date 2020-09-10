package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"

	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
	request "github.com/williamchang80/sea-apd/dto/request/admin"
	user_mock_repository "github.com/williamchang80/sea-apd/mocks/repository/user"
	user_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/user"
	"github.com/williamchang80/sea-apd/usecase/user"
)

var (
	mockData = admin.Admin{
		Name:     "admin",
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

func TestAdminController_RegisterAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request request.Admin
	}
	defer ctrl.Finish()

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.AdminUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockData,
			},
			wantErr: false,
			initMock: func() domain.AdminUsecase {
				c := user_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: request.Admin{},
			},
			wantErr: false,
			initMock: func() domain.AdminUsecase {
				c := user_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.POST, "api/users/register-admin", strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("RegisterADmin() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewAdminController(c, mock)
			if got := controller.RegisterAdmin(ctx); (got != nil) != tt.wantErr {
				t.Errorf("RegisterAdmin() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}
