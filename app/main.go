package main

import (
	"net/http"

	"github.com/44taka/twitter-trends-api/infrastructure"
	"github.com/44taka/twitter-trends-api/infrastructure/persistence"
	"github.com/44taka/twitter-trends-api/presentation/handler"
	"github.com/44taka/twitter-trends-api/usecase"
	"github.com/gin-contrib/cors"
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

	twitterTrendPersistence := persistence.NewTwitterTrendPersistence(db.Connect())
	twitterTrendUseCase := usecase.NewTwitterTrendUseCase(twitterTrendPersistence)
	twitterTrendHandler := handler.NewTwitterTrendHandler(twitterTrendUseCase)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// r.GET("/twitter/trends", func(ctx *gin.Context) { twitterTrendHandler.FindAll(ctx) })
	r.GET("/twitter/trends", func(ctx *gin.Context) { twitterTrendHandler.Find(ctx) })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Routing.Port)
}
