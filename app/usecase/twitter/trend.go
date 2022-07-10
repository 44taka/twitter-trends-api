package usecase

import (
	"fmt"
	"time"

	model "github.com/44taka/twitter-trends-api/domain/model/twitter"
	repository "github.com/44taka/twitter-trends-api/domain/repository/twitter"
	"github.com/gin-gonic/gin"
)

type TwitterTrendUseCase interface {
	FindAll(ctx *gin.Context) ([]*model.TwitterTrend, error)
	Find(ctx *gin.Context) (*model.TwitterTrendResult, error)
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

func (ttu twitterTrendUseCase) Find(ctx *gin.Context) (twitter_trends *model.TwitterTrendResult, err error) {
	// nowDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.UTC)
	// nowDate := time.Date(time.Now().Year(), time.Now().Month(), 9, 10, 0, 0, 0, time.UTC)
	// startDateTime := nowDate.Add(-(0) * time.Hour)
	// endDateTime := startDateTime.Add(1 * time.Hour)
	// twitter_trends, err = ttu.twitterTrendRepository.Find(ctx, startDateTime, endDateTime)
	// if err != nil {
	// 	return twitter_trends, err
	// }
	// return twitter_trends[:5], nil

	// var twitter_trends *model.TwitterTrendResult
	nowDate := time.Date(time.Now().Year(), time.Now().Month(), 9, 10, 0, 0, 0, time.UTC)
	for _, n := range [...]int{0, 1} {
		startDateTime := nowDate.Add(-(time.Duration(n)) * time.Hour)
		endDateTime := startDateTime.Add(1 * time.Hour)
		fmt.Printf(endDateTime.String())
		r, _ := ttu.twitterTrendRepository.Find(ctx, startDateTime, endDateTime)
		// if err != nil {
		// 	return twitter_trends, err
		// }
		twitter_trends.Result = append(twitter_trends.Result, r)
	}
	return twitter_trends, nil
}
