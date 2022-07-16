package usecase

import (
	"strconv"
	"time"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/domain/repository"
	"github.com/gin-gonic/gin"
)

type TwitterTrendResponse struct {
	Message string               `json:"message"`
	Result  []TwitterTrendResult `json:"result"`
}

type TwitterTrendResult struct {
	Label  string                `json:"label"`
	Trends []*model.TwitterTrend `json:"trends"`
}

type TwitterTrendUseCase interface {
	Find(ctx *gin.Context) (TwitterTrendResponse, error)
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

func (ttu twitterTrendUseCase) Find(ctx *gin.Context) (TwitterTrendResponse, error) {
	var twitter_trend_response TwitterTrendResponse

	nowDate := time.Date(time.Now().Year(), time.Now().Month(), 9, 10, 0, 0, 0, time.UTC)
	for _, n := range [...]int{0, 1, 3} {
		var twitter_trend_result TwitterTrendResult

		startDateTime := nowDate.Add(-(time.Duration(n)) * time.Hour)
		endDateTime := startDateTime.Add(1 * time.Hour)
		twitter_trends, err := ttu.twitterTrendRepository.Find(ctx, startDateTime, endDateTime)
		if err != nil {
			return twitter_trend_response, err
		}

		if n == 0 {
			twitter_trend_result.Label = "現在"
		} else {
			twitter_trend_result.Label = strconv.Itoa(n) + "時間前"
		}
		twitter_trend_result.Trends = twitter_trends
		twitter_trend_response.Result = append(twitter_trend_response.Result, twitter_trend_result)
	}

	return twitter_trend_response, nil
}

func (ttu twitterTrendUseCase) FindAll(ctx *gin.Context) (twitter_trends []*model.TwitterTrend, err error) {
	twitter_trends, err = ttu.twitterTrendRepository.FindAll(ctx)
	if err != nil {
		return twitter_trends, err
	}
	return twitter_trends, nil
}
