package service_user

import (
	"wire/internal/common/res"
	"wire/internal/global"
	"wire/internal/repo/repo_user"
	"wire/internal/type/type_user"
	"wire/internal/utils"
	"wire/internal/utils/snowflake"
)

type IUserService interface {
	GetUserInfo(req type_user.UserInfoRequest) (resp type_user.UserInfoResponse, err error)
	CreateUser(req type_user.UserCreateRequest) (resp type_user.UserCreateResponse, err error)
	UpdateUser(req type_user.UserUpdateRequest) (err error)
	DeleteUser(req type_user.UserDeleteRequest) (err error)
}
type UserService struct {
	UserRepo repo_user.IUserRepo
}

// 检查结构体是否实现了某接口的防御代码
var _ IUserService = (*UserService)(nil)

func (s *UserService) GetUserInfo(req type_user.UserInfoRequest) (resp type_user.UserInfoResponse, err error) {
	data, err := s.UserRepo.GetUserInfo(utils.StringToInt64(req.ID))
	if err != nil {
		return type_user.UserInfoResponse{}, res.ErrResp(err, res.USER_NOT_EXIST)
	}
	resp = type_user.UserInfoResponse{
		Username: data.Username,
		Password: data.Password,
	}
	return
}

func (s *UserService) CreateUser(req type_user.UserCreateRequest) (resp type_user.UserCreateResponse, err error) {
	id := snowflake.GetIntId(global.Node)
	err = s.UserRepo.CreateUser(id, req)
	if err != nil {
		return type_user.UserCreateResponse{}, res.ErrResp(err, res.USER_CREATE_ERROR)
	}
	resp.ID = id
	return
}

func (s *UserService) UpdateUser(req type_user.UserUpdateRequest) (err error) {
	id := utils.StringToInt64(req.ID)
	_, err = s.UserRepo.GetUserInfo(id)
	if err != nil {
		return res.ErrResp(err, res.USER_NOT_EXIST)
	}
	err = s.UserRepo.UpdateUser(id, req.Username, req.Password)
	if err != nil {
		return res.ErrResp(err, res.USER_UPDATE_ERROR)
	}
	return
}

func (s *UserService) DeleteUser(req type_user.UserDeleteRequest) (err error) {
	err = s.UserRepo.DeleteUser(utils.StringToInt64(req.ID))
	if err != nil {
		return res.ErrResp(err, res.USER_DELETE_ERROR)
	}
	return
}
