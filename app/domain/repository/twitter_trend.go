package repository

import (
	"time"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/gin-gonic/gin"
)

type TwitterTrendRepository interface {
	Find(ctx *gin.Context, startDateTime time.Time, endDateTime time.Time) ([]*model.TwitterTrend, error)
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
}
