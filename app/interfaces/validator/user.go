package validator

// formのバリデーション定義
type UserForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// uriのバリデーション定義
type UserUri struct {
	ID int `uri:"id" binding:"required,numeric"`
}
