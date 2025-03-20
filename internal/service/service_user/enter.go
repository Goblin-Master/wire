package service_user

import (
	"wire/internal/repo/repo_user"
	"wire/internal/type/type_user"
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

func (s *UserService) GetUserInfo(req type_user.UserInfoRequest) (resp type_user.UserInfoResponse, err error) {
	resp, err = s.UserRepo.GetUserInfo(req)
	return
}

func (s *UserService) CreateUser(req type_user.UserCreateRequest) (resp type_user.UserCreateResponse, err error) {
	resp, err = s.UserRepo.CreateUser(req)
	return
}

func (s *UserService) UpdateUser(req type_user.UserUpdateRequest) (err error) {
	return s.UserRepo.UpdateUser(req)
}

func (s *UserService) DeleteUser(req type_user.UserDeleteRequest) (err error) {
	return s.UserRepo.DeleteUser(req)
}
