package handler

import (
	"net/http"

	"github.com/44taka/twitter-trends/interfaces/validator"
	"github.com/44taka/twitter-trends/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) FindAll(ctx *gin.Context) {
	// ユースケース呼び出し
	users, err := uh.userUseCase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found user",
			"result":  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user all",
		"result":  users,
	})
	return
}

func (uh userHandler) FindById(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
			"result":  nil,
		})
		return
	}
	user, err := uh.userUseCase.FindById(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"result":  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user",
		"result":  user,
	})
	return
}

func (uh userHandler) Create(ctx *gin.Context) {
	var form validator.UserForm
	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
			"result":  nil,
		})
		return
	}
	err = uh.userUseCase.Create(ctx, form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
			"result":  nil,
		})
		return
	}
	// ctx.Writer.WriteHeader(http.StatusCreated)
	ctx.AbortWithStatus(http.StatusCreated)
	return
}

func (uh userHandler) Update(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
			"result":  nil,
		})
		return
	}
	var form validator.UserForm
	err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
			"result":  nil,
		})
		return
	}
	err = uh.userUseCase.Update(ctx, uri.ID, form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
			"result":  nil,
		})
		return
	}
	ctx.AbortWithStatus(http.StatusNoContent)
	return
}

func (uh userHandler) Delete(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
			"result":  nil,
		})
		return
	}
	err = uh.userUseCase.Delete(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
			"result":  nil,
		})
		return
	}
	ctx.AbortWithStatus(http.StatusNoContent)
	return
}
