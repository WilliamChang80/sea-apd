package user

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/admin"
	user_repo "github.com/williamchang80/sea-apd/mocks/repository/user"
)

func TestNewAdminUseCase(t *testing.T) {
	type args struct {
		repository user.UserRepository
	}
	tests := []struct {
		name string
		args args
		want user.AdminUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: &AdminUsecase{
				ur: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdminUseCase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAdminUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToDomain(t *testing.T) {
	type args struct {
		adminReq request.Admin
	}
	tests := []struct {
		name string
		args args
		want user.User
	}{
		{
			name: "success",
			args: args{
				adminReq: request.Admin{
					Name:     "name",
					Email:    "email@email.com",
					Password: "password",
				},
			},
			want: user.User{
				Name:     "name",
				Email:    "email@email.com",
				Password: "password",
				Role:     "0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToDomain(tt.args.adminReq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToDomain() - %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminUsecase_RegisterAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		adminReq request.Admin
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user.AdminUsecase
	}{
		{
			name: "failed cause duplicate",
			args: args{
				adminReq: request.Admin{
					Name:     "name",
					Email:    "email@email.net",
					Password: "password",
				},
			},
			wantErr: true,
			initMock: func() user.AdminUsecase {
				r := user_repo.NewMockRepository(ctrl)
				return NewAdminUseCase(r)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				adminReq: request.Admin{},
			},
			wantErr: true,
			initMock: func() user.AdminUsecase {
				r := user_repo.NewMockRepository(ctrl)
				return NewAdminUseCase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.RegisterAdmin(tt.args.adminReq)
			if err != nil && !tt.wantErr {
				t.Errorf("AdminUseCase.RegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
