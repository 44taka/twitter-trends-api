package usecase

import (
	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/domain/repository"
	"github.com/gin-gonic/gin"
)

type WoeidUseCase interface {
	FindAll(ctx *gin.Context) ([]*model.Woeid, error)
}

type woeidUseCase struct {
	woeidRepository repository.WoeidRepository
}

func NewWoeidUseCase(wr repository.WoeidRepository) WoeidUseCase {
	return &woeidUseCase{
		woeidRepository: wr,
	}
}

func (wu woeidUseCase) FindAll(ctx *gin.Context) (woeid []*model.Woeid, err error) {
	woeid, err = wu.woeidRepository.FindAll(ctx)
	if err != nil {
		return woeid, err
	}
	return woeid, nil
}
