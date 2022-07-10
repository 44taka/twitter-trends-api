package repository

import (
	"time"

	model "github.com/44taka/twitter-trends-api/domain/model/twitter"
	"github.com/gin-gonic/gin"
)

type TwitterTrendRepository interface {
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
	Find(ctx *gin.Context, startTime time.Time, endTime time.Time) ([]*model.TwitterTrend, error)
}
