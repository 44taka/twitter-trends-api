package persistence

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// TODO::テストデータのリファクタやる
	var users []*model.UserTestData
	testdata, _ := ioutil.ReadFile("../../testdata/user_fixture.json")
	err := json.Unmarshal(testdata, &users)
	if err != nil {
		os.Exit(1)
	}
	db := infrastructure.NewDB(infrastructure.NewConfig()).Connect()
	db.Exec("delete from users")
	for i := 0; i < len(users); i++ {
		db.Exec(
			"insert into users values (?, ?, ?)",
			users[i].ID,
			users[i].Name,
			users[i].Password,
		)
	}
	status := m.Run()
	// db.Exec("delete from users")
	os.Exit(status)
}

func TestUserPersistenceFindByAll(t *testing.T) {
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
				{
					ID:       4,
					Name:     "shiro",
					Password: "password",
				},
				{
					ID:       5,
					Name:     "goro",
					Password: "password",
				},
			},
			err: nil,
		},
		// TODO:データなしのテストパターンをどうする？
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// TODO:gormはモック化したい
			db := infrastructure.NewDB(infrastructure.NewConfig())
			userPersistence := NewUserPersistence(db.Connection)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			got, err := userPersistence.FindAll(ctx)

			asserts := assert.New(t)
			if err != nil {
				asserts.EqualError(tt.err, err.Error())
			}
			asserts.Equal(tt.want, got)
		})
	}
}

func TestUserPersistenceFindById(t *testing.T) {
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
			id:   999,
			want: model.User{},
			err:  errors.New("user is not found"),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// TODO:gormはモック化したい
			config := infrastructure.NewConfig()
			db := infrastructure.NewDB(config)
			userPersistence := NewUserPersistence(db.Connection)

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			got, err := userPersistence.FindById(ctx, tt.id)

			asserts := assert.New(t)
			if err != nil {
				asserts.EqualError(tt.err, err.Error())
			}
			asserts.Equal(tt.want, got)
		})
	}
}

func TestUserPersistenceCreate(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ登録",
			want: model.User{
				ID:       6,
				Name:     "rokuro",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "エラー",
			want: model.User{
				Name:     "ichiro",
				Password: "passwordaaaaaaaaaaaaaaaaaaaaaa",
			},
			err: errors.New("failed create user..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// TODO:gormはモック化したい
			config := infrastructure.NewConfig()
			db := infrastructure.NewDB(config)
			userPersistence := NewUserPersistence(db.Connection)

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := userPersistence.Create(ctx, tt.want.Name, tt.want.Password)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(tt.err, err.Error())
			} else {
				asserts.NoError(err)
			}
		})
	}
}

func TestUserPersistenceUpdate(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ更新",
			want: model.User{
				ID:       1,
				Name:     "ichiro_edit",
				Password: "passss",
			},
			err: nil,
		},
		{
			name: "データなし",
			want: model.User{ID: 999},
			err:  errors.New("user is not found"),
		},
		{
			name: "エラー",
			want: model.User{
				ID:       5,
				Name:     "goro",
				Password: "passwordaaaaaaaaaaaaaaaaaaaaaa",
			},
			err: errors.New("failed update user..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// TODO:gormはモック化したい
			db := infrastructure.NewDB(infrastructure.NewConfig())
			userPersistence := NewUserPersistence(db.Connection)

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := userPersistence.Update(ctx, tt.want.ID, tt.want.Name, tt.want.Password)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(tt.err, err.Error())
			} else {
				asserts.NoError(err)
			}
		})
	}
}

func TestUserPersistenceDelete(t *testing.T) {
	cases := []struct {
		name string
		want model.User
		err  error
	}{
		{
			name: "データ削除",
			want: model.User{ID: 5},
			err:  nil,
		},
		{
			name: "データなし",
			want: model.User{ID: 999},
			err:  errors.New("user is not found"),
		},
		{
			name: "エラー",
			want: model.User{ID: 999999999999999999},
			err:  errors.New("failed delete user..."),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// TODO:gormはモック化したい
			db := infrastructure.NewDB(infrastructure.NewConfig())
			userPersistence := NewUserPersistence(db.Connection)

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := userPersistence.Delete(ctx, tt.want.ID)

			asserts := assert.New(t)
			if err != nil {
				asserts.Error(err)
				asserts.EqualError(tt.err, err.Error())
			} else {
				asserts.NoError(err)
			}
		})
	}
}
