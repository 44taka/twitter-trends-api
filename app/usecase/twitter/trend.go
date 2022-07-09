package usecase

import (
	model "github.com/44taka/twitter-trends-api/domain/model/twitter"
	repository "github.com/44taka/twitter-trends-api/domain/repository/twitter"
	"github.com/gin-gonic/gin"
)

type TwitterTrendUseCase interface {
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
}

type twitterTrendUseCase struct {
	twitterTrendRepository repository.TwitterTrendRepository
}

func NewTwitterTrendUseCase(ttr repository.TwitterTrendRepository) TwitterTrendUseCase {
	return &twitterTrendUseCase{
		twitterTrendRepository: ttr,
	}
}

func (ttu twitterTrendUseCase) FindAll(ctx *gin.Context) (twitter_trends []*model.TwitterTrend, err error) {
	twitter_trends, err = ttu.twitterTrendRepository.FindAll(ctx)
	if err != nil {
		return twitter_trends, err
	}
	return twitter_trends, nil
}
