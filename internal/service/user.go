package service

import (
	"wire/internal/model"
	"wire/internal/repo"
)

type IUserService interface {
	GetUserInfo() *model.User
}
type UserService struct {
	UserRepo repo.IUserRepo
}

func (s *UserService) GetUserInfo() *model.User {
	userInfo := s.UserRepo.GetUserInfo()
	return userInfo
}
