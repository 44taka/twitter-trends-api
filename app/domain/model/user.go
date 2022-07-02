package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"-" form:"password" binding:"required"`
}

type UserUri struct {
	ID int `uri:"id" binding:"required,numeric"`
}

type UserTestData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
