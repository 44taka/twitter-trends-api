package main

import (
	"github.com/44taka/twitter-trends-api/infrastructure"
	persistence_twitter "github.com/44taka/twitter-trends-api/infrastructure/persistence/twitter"
	handler_twitter "github.com/44taka/twitter-trends-api/interfaces/handler/twitter"
	usecase_twitter "github.com/44taka/twitter-trends-api/usecase/twitter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// コンフィグ読み込み
	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)

	twitterTrendPersistence := persistence_twitter.NewTwitterTrendPersistence(db.Connect())
	twitterTrendUseCase := usecase_twitter.NewTwitterTrendUseCase(twitterTrendPersistence)
	twitterTrendHandler := handler_twitter.NewTwitterTrendHandler(twitterTrendUseCase)

	r := gin.Default()

	// r.GET("/twitter/trends", func(ctx *gin.Context) { twitterTrendHandler.FindAll(ctx) })
	r.GET("/twitter/trends", func(ctx *gin.Context) { twitterTrendHandler.Find(ctx) })

	// MEMO: CRUDのエンドポイントのサンプルとして残しておく
	// ハンドラー読み込み
	// userPersistence := persistence.NewUserPersistence(db.Connect())
	// userUseCase := usecase.NewUserUseCase(userPersistence)
	// userHandler := handler.NewUserHandler(userUseCase)
	// r.GET("/users", func(ctx *gin.Context) { userHandler.FindAll(ctx) })
	// r.GET("/users/:id", func(ctx *gin.Context) { userHandler.FindById(ctx) })
	// r.POST("/users", func(ctx *gin.Context) { userHandler.Create(ctx) })
	// r.PUT("/users/:id", func(ctx *gin.Context) { userHandler.Update(ctx) })
	// r.DELETE("/users/:id", func(ctx *gin.Context) { userHandler.Delete(ctx) })

	r.Run(":" + config.Routing.Port)
}
