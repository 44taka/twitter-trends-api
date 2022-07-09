package repository

import (
	"github.com/44taka/twitter-trends-api/domain/model"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	FindAll(ctx *gin.Context) ([]*model.User, error)
	FindById(ctx *gin.Context, id int) (model.User, error)
	Create(ctx *gin.Context, name string, password string) error
	Update(ctx *gin.Context, id int, name string, password string) error
	Delete(ctx *gin.Context, id int) error
}
