package user

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	domain "github.com/williamchang80/sea-apd/domain/user"

	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
)

func TestNewUserRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.UserRepository
	}{
		{
			name: "success with null value on db",
			args: args{
				db: nil,
			},
			want: &UserRepository{
				db: nil,
			},
		},
		{
			name: "Success with value on db",
			args: args{
				db: db,
			},
			want: &UserRepository{
				db: db,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		user domain.User
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "fail with invalid db query",
			args: args{
				user: domain.User{
					Username: "username",
					Email:    "email",
					Password: "password",
					Role:     "0",
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectQuery(`
				"INSERT INTO "users"
					("created_at",
					"updated_at",
					"deleted_at",
					"username",
					"email",
					"password",
					"role")
				VALUES ($1,$2,$3,$4,$5,$6,$7)
				RETURNING "users"."id"
				`).WithArgs(
					time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC),
					time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC),
					nil,
					nil,
					"email@email.com",
					nil,
					"0",
				).WillReturnRows(sqlmock.NewRows([]string{
					"id",
				}))
				mocks.ExpectCommit()
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := UserRepository{
				db: tt.initMock(),
			}
			err := ur.CreateUser(tt.args.user)
			if err != nil && !tt.wantErr {
				t.Errorf("UserRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
