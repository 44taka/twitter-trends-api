package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/44taka/twitter-trends/domain/model"
	"github.com/44taka/twitter-trends/infrastructure"
	"github.com/44taka/twitter-trends/infrastructure/persistence"
	"github.com/44taka/twitter-trends/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// TODO:テストデータのリファクタやる
	var users []*model.UserTestData
	testdata, _ := ioutil.ReadFile("../../testdata/user_fixture.json")
	err := json.Unmarshal(testdata, &users)
	if err != nil {
		os.Exit(1)
	}
	db := infrastructure.NewDB(infrastructure.NewTestConfig()).Connect()
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

func TestFindAll(t *testing.T) {
	type expectedResponse struct {
		Message string
		Result  []model.User
	}
	cases := []struct {
		name string
		want struct {
			Message string
			Result  []model.User
		}
		wantStatus int
	}{
		{
			name: "データ取得",
			want: struct {
				Message string
				Result  []model.User
			}{
				Message: "get user all",
				Result: []model.User{
					{
						ID:   1,
						Name: "ichiro",
					},
					{
						ID:   2,
						Name: "jiro",
					},
					{
						ID:   3,
						Name: "saburo",
					},
					{
						ID:   4,
						Name: "shiro",
					},
					{
						ID:   5,
						Name: "goro",
					},
				},
			},
			wantStatus: http.StatusOK,
		},
		// TODO:データ0件のテストパターン考える
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)

			userHandler := getHandler()
			userHandler.FindAll(ctx)

			var expectedResponse expectedResponse
			_ = json.Unmarshal(response.Body.Bytes(), &expectedResponse)
			asserts := assert.New(t)
			asserts.Equal(tt.wantStatus, response.Result().StatusCode)
			asserts.Equal(tt.want.Message, expectedResponse.Message)
			asserts.Equal(tt.want.Result, expectedResponse.Result)
		})
	}
}

func TestFindById(t *testing.T) {
	type expectedResponse struct {
		Message string
		Result  model.User
	}
	cases := []struct {
		name string
		id   int
		want struct {
			Message string
			Result  model.User
		}
		wantStatus int
	}{
		{
			name: "データ取得",
			id:   1,
			want: struct {
				Message string
				Result  model.User
			}{
				Message: "get user",
				Result: model.User{
					ID:   1,
					Name: "ichiro",
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "データなし",
			id:   999,
			want: struct {
				Message string
				Result  model.User
			}{
				Message: "user is not found",
				Result:  model.User{},
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(tt.id), nil)
			ctx.Params = gin.Params{
				gin.Param{"id", strconv.Itoa(tt.id)},
			}

			userHandler := getHandler()
			userHandler.FindById(ctx)

			var expectedResponse expectedResponse
			_ = json.Unmarshal(response.Body.Bytes(), &expectedResponse)
			asserts := assert.New(t)
			asserts.Equal(tt.wantStatus, response.Result().StatusCode)
			asserts.Equal(tt.want.Message, expectedResponse.Message)
			asserts.Equal(tt.want.Result, expectedResponse.Result)
		})
	}
}

func TestCreate(t *testing.T) {
	cases := []struct {
		name       string
		data       model.User
		wantStatus int
	}{
		{
			name: "データ登録",
			data: model.User{
				Name:     "tojiro",
				Password: "password",
			},
			wantStatus: http.StatusCreated,
		},
		// TODO:バリデーションエラーのテストケースも
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			requestValues := url.Values{}
			requestValues.Set("name", tt.data.Name)
			requestValues.Set("password", tt.data.Password)

			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/users",
				strings.NewReader(requestValues.Encode()),
			)
			ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// TODO::リファクタしたい
			userHandler := getHandler()
			userHandler.Create(ctx)

			asserts := assert.New(t)
			asserts.Equal(tt.wantStatus, response.Result().StatusCode)
		})
	}
}

func TestUpdate(t *testing.T) {
	cases := []struct {
		name       string
		data       model.User
		wantStatus int
	}{
		{
			name: "データ更新",
			data: model.User{
				ID:       5,
				Name:     "tojiro",
				Password: "password",
			},
			wantStatus: http.StatusNoContent,
		},
		// TODO:バリデーションエラーのテストケースも
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			requestValues := url.Values{}
			requestValues.Set("name", tt.data.Name)
			requestValues.Set("password", tt.data.Password)

			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request, _ = http.NewRequest(
				http.MethodPut,
				"/users/"+strconv.Itoa(tt.data.ID),
				strings.NewReader(requestValues.Encode()),
			)
			ctx.Params = gin.Params{
				gin.Param{"id", strconv.Itoa(tt.data.ID)},
			}
			ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// TODO::リファクタしたい
			userHandler := getHandler()
			userHandler.Update(ctx)

			asserts := assert.New(t)
			asserts.Equal(tt.wantStatus, response.Result().StatusCode)
		})
	}
}

func TestDelete(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		data       model.User
		wantStatus int
	}{
		{
			name:       "データ削除",
			id:         5,
			wantStatus: http.StatusNoContent,
		},
		// TODO:バリデーションエラーのテストケースも
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request, _ = http.NewRequest(
				http.MethodDelete,
				"/users/"+strconv.Itoa(tt.id),
				nil,
			)
			ctx.Params = gin.Params{
				gin.Param{"id", strconv.Itoa(tt.id)},
			}

			// TODO::リファクタしたい
			userHandler := getHandler()
			userHandler.Delete(ctx)

			asserts := assert.New(t)
			asserts.Equal(tt.wantStatus, response.Result().StatusCode)
		})
	}
}

func getHandler() UserHandler {
	config := infrastructure.NewTestConfig()
	db := infrastructure.NewDB(config)
	userPersistence := persistence.NewUserPersistence(db.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := NewUserHandler(userUseCase)
	return userHandler
}
