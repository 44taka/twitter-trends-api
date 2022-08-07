package persistence

import (
	"errors"
	"time"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type twitterTrendPersistence struct {
	conn *gorm.DB
}

func NewTwitterTrendPersistence(conn *gorm.DB) repository.TwitterTrendRepository {
	return &twitterTrendPersistence{conn: conn}
}

func (ttp twitterTrendPersistence) Find(ctx *gin.Context, startDateTime time.Time, endDateTime time.Time) ([]*model.TwitterTrend, error) {
	const layout = "2006-01-02 15:04:05"
	twitter_trends := []*model.TwitterTrend{}
	r := ttp.conn.
		Where("created_at >= ?", startDateTime.Format(layout)).
		Where("created_at < ?", endDateTime.Format(layout)).
		Order("created_at desc").
		Order("rank").
		Limit(50).
		Find(&twitter_trends)
	if r.Error != nil {
		return twitter_trends, errors.New("twitter trends are not found")
	}
	return twitter_trends, nil
}

func (ttp twitterTrendPersistence) FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error) {
	twitter_trends := []*model.TwitterTrend{}
	r := ttp.conn.
		Order("created_at desc").
		Order("rank").
		Limit(50).
		Find(&twitter_trends)
	if r.Error != nil {
		return twitter_trends, errors.New("twitter trends are not found")
	}
	return twitter_trends, nil
}
