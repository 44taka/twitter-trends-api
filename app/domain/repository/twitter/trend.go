package repository

import (
	model "github.com/44taka/twitter-trends-api/domain/model/twitter"
	"github.com/gin-gonic/gin"
)

type TwitterTrendRepository interface {
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
}
