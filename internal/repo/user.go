package repo

import (
	"gorm.io/gorm"
	"wire/internal/model"
)

type IUserRepo interface {
	GetUserInfo() *model.User
}

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) GetUserInfo() *model.User {
	var user model.User
	r.DB.Take(&user, 1)
	return &user
}
