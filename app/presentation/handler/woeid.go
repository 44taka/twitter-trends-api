package handler

import (
	"net/http"

	"github.com/44taka/twitter-trends-api/usecase"
	"github.com/gin-gonic/gin"
)

type WoeidHandler interface {
	FindAll(ctx *gin.Context)
}

type woeidHandler struct {
	woeidUseCase usecase.WoeidUseCase
}

func NewWoeidHandler(wu usecase.WoeidUseCase) WoeidHandler {
	return &woeidHandler{
		woeidUseCase: wu,
	}
}

func (wh woeidHandler) FindAll(ctx *gin.Context) {
	woeid, err := wh.woeidUseCase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, woeid)
		return
	}
	ctx.JSON(http.StatusOK, woeid)
	return
}
