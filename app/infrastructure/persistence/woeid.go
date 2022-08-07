package persistence

import (
	"errors"

	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/44taka/twitter-trends-api/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type woeidPersistence struct {
	conn *gorm.DB
}

func NewWoeidPersistence(conn *gorm.DB) repository.WoeidRepository {
	return &woeidPersistence{conn: conn}
}

func (wp woeidPersistence) FindAll(ctx *gin.Context) ([]*model.Woeid, error) {
	woeid := []*model.Woeid{}
	r := wp.conn.Table("woeid").Find(&woeid)
	if r.Error != nil {
		return woeid, errors.New("woeid are not found")
	}
	return woeid, nil
}
