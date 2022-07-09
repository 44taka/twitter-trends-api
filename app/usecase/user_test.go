package usecase

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestUseCaseUserFindAll...
func TestUseCaseUserFindAll(t *testing.T) {
	cases := []struct {
		name string
		want []*model.User
		err  error
	}{
		{
			name: "データ取得",
			want: []*model.User{
				{
					ID:       1,
					Name:     "ichiro",
					Password: "password",
				},
				{
					ID:       2,
					Name:     "jiro",
					Password: "password",
				},
				{
					ID:       3,
					Name:     "saburo",
					Password: "password",
				},
			},
			err: nil,
		},
		{
			name: "データなし",
			want: []*model.User{},
			err:  nil,
		},
		{
			name: "エラー",
			want: []*model.User{},
			err:  errors.New("error..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().FindAll(ctx).Return(tt.want, tt.err)

			userUseCase := NewUserUseCase(mockUserRepository)
			got, err := userUseCase.FindAll(ctx)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(err, "error...")
			}
			asserts.Equal(tt.want, got)
		})
	}
}

// TestUseCaseUserFindById...
func TestUseCaseUserFindById(t *testing.T) {
	cases := []struct {
		name string
		id   int
		want model.User
		err  error
	}{
		{
			name: "データ取得",
			id:   1,
			want: model.User{
				ID:       1,
				Name:     "ichiro",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "データなし",
			id:   2,
			want: model.User{},
			err:  nil,
		},
		{
			name: "エラー",
			id:   3,
			want: model.User{},
			err:  errors.New("error..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().FindById(ctx, tt.id).Return(tt.want, tt.err)

			userUseCase := NewUserUseCase(mockUserRepository)
			got, err := userUseCase.FindById(ctx, tt.id)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(err, "error...")
			}
			asserts.Equal(tt.want, got)
		})
	}
}

// TestUseCaseUserCreate...
func TestUseCaseUserCreate(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ登録",
			want: model.User{
				Name:     "ichiro",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "エラー",
			want: model.User{
				Name:     "jiro",
				Password: "password",
			},
			err: errors.New("error..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().Create(ctx, tt.want.Name, tt.want.Password).Return(tt.err)

			userUseCase := NewUserUseCase(mockUserRepository)
			err := userUseCase.Create(ctx, tt.want.Name, tt.want.Password)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(err, "error...")
			} else {
				asserts.NoError(err)
			}
		})
	}
}

// TestUseCaseUserUpdate...
func TestUseCaseUserUpdate(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ更新",
			want: model.User{
				ID:       1,
				Name:     "ichiro",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "エラー",
			want: model.User{
				ID:       2,
				Name:     "jiro",
				Password: "password",
			},
			err: errors.New("error..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().Update(ctx, tt.want.ID, tt.want.Name, tt.want.Password).Return(tt.err)

			userUseCase := NewUserUseCase(mockUserRepository)
			err := userUseCase.Update(ctx, tt.want.ID, tt.want.Name, tt.want.Password)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(err, "error...")
			} else {
				asserts.NoError(err)
			}
		})
	}
}

// TestUseCaseUserDelete...
func TestUseCaseUserDelete(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ削除",
			want: model.User{
				ID:       1,
				Name:     "ichiro",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "エラー",
			want: model.User{
				ID:       2,
				Name:     "jiro",
				Password: "password",
			},
			err: errors.New("error..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().Delete(ctx, tt.want.ID).Return(tt.err)

			userUseCase := NewUserUseCase(mockUserRepository)
			err := userUseCase.Delete(ctx, tt.want.ID)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(err, "error...")
			} else {
				asserts.NoError(err)
			}
		})
	}
}
