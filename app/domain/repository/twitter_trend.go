package repository

import (
	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/gin-gonic/gin"
)

type TwitterTrendRepository interface {
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
}
