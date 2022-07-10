package handler

import (
	"net/http"

	"github.com/44taka/twitter-trends-api/usecase"
	"github.com/gin-gonic/gin"
)

type TwitterTrendHandler interface {
	FindAll(ctx *gin.Context)
}

type twitterTrendHandler struct {
	twitterTrendUseCase usecase.TwitterTrendUseCase
}

func NewTwitterTrendHandler(ttu usecase.TwitterTrendUseCase) TwitterTrendHandler {
	return &twitterTrendHandler{
		twitterTrendUseCase: ttu,
	}
}

func (tth twitterTrendHandler) FindAll(ctx *gin.Context) {
	twitter_trends, err := tth.twitterTrendUseCase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found user",
			"result":  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get twitter trends all",
		"result":  twitter_trends,
	})
	return
}
