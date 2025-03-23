package repo_user

import (
	"gorm.io/gorm"
	"wire/internal/model"
	"wire/internal/type/type_user"
)

type IUserRepo interface {
	GetUserInfo(id int64) (model model.User, err error)
	CreateUser(id int64, req type_user.UserCreateRequest) (err error)
	UpdateUser(id int64, name, password string) (err error)
	DeleteUser(id int64) (err error)
}

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) GetUserInfo(id int64) (model model.User, err error) {
	err = r.DB.Take(&model, id).Error
	return
}

func (r *UserRepo) CreateUser(id int64, req type_user.UserCreateRequest) (err error) {
	return r.DB.Create(&model.User{
		ID:       id,
		Username: req.Username,
		Password: req.Password,
	}).Error
}

func (r *UserRepo) UpdateUser(id int64, name, password string) (err error) {
	return r.DB.Model(&model.User{}).Where("id = ?", id).Updates(map[string]any{
		"username": name,
		"password": password,
	}).Error
}

func (r *UserRepo) DeleteUser(id int64) (err error) {
	return r.DB.Delete(model.User{}, id).Error
}
