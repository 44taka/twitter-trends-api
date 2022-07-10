package usecase

import (
	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/domain/repository"
	"github.com/gin-gonic/gin"
)

type TwitterTrendResult struct {
	Message string                `json:"message"`
	Result  []*model.TwitterTrend `json:"result"`
}

type TwitterTrendUseCase interface {
	Find(ctx *gin.Context) (TwitterTrendResult, error)
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

func (ttu twitterTrendUseCase) Find(ctx *gin.Context) (TwitterTrendResult, error) {
	var twitter_trend_result TwitterTrendResult
	twitter_trends, err := ttu.twitterTrendRepository.FindAll(ctx)
	if err != nil {
		return twitter_trend_result, err
	}
	twitter_trend_result.Message = "test!!"
	twitter_trend_result.Result = twitter_trends

	return twitter_trend_result, nil
}

func (ttu twitterTrendUseCase) FindAll(ctx *gin.Context) (twitter_trends []*model.TwitterTrend, err error) {
	twitter_trends, err = ttu.twitterTrendRepository.FindAll(ctx)
	if err != nil {
		return twitter_trends, err
	}
	return twitter_trends, nil
}
