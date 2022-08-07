package repository

import (
	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/gin-gonic/gin"
)

type WoeidRepository interface {
	FindAll(ctx *gin.Context) ([]*model.Woeid, error)
}
