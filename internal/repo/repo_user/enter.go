package repo_user

import (
	"gorm.io/gorm"
	"wire/internal/global"
	"wire/internal/model"
	"wire/internal/type/type_user"
	"wire/internal/utils"
	"wire/internal/utils/snowflake"
)

type IUserRepo interface {
	GetUserInfo(req type_user.UserInfoRequest) (resp type_user.UserInfoResponse, err error)
	CreateUser(req type_user.UserCreateRequest) (resp type_user.UserCreateResponse, err error)
	UpdateUser(req type_user.UserUpdateRequest) (err error)
	DeleteUser(req type_user.UserDeleteRequest) (err error)
}

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) GetUserInfo(req type_user.UserInfoRequest) (resp type_user.UserInfoResponse, err error) {
	var user model.User
	err = r.DB.Take(user, utils.StringToInt64(req.ID)).Error
	if err != nil {
		return
	}
	resp = type_user.UserInfoResponse{
		Username: user.Username,
		Password: user.Password,
	}
	return
}

func (r *UserRepo) CreateUser(req type_user.UserCreateRequest) (resp type_user.UserCreateResponse, err error) {
	resp.ID = snowflake.GetIntId(global.Node)
	err = r.DB.Create(&model.User{
		ID:       resp.ID,
		Username: req.Username,
		Password: req.Password,
	}).Error
	return
}

func (r *UserRepo) UpdateUser(req type_user.UserUpdateRequest) (err error) {
	return r.DB.Model(&model.User{}).Updates(req).Error
}

func (r *UserRepo) DeleteUser(req type_user.UserDeleteRequest) (err error) {
	return r.DB.Delete(model.User{}, req.ID).Error
}
