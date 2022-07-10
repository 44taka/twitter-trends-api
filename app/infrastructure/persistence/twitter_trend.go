package persistence

import (
	"errors"

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

func (ttp twitterTrendPersistence) FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error) {
	twitter_trends := []*model.TwitterTrend{}
	r := ttp.conn.Find(&twitter_trends)
	if r.Error != nil {
		return twitter_trends, errors.New("twitter trends are not found")
	}
	return twitter_trends[:10], nil
}
