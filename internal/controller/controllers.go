package controller

import (
	"github.com/google/wire"
	"wire/internal/controller/controller_user"
	"wire/internal/repo/repo_user"
	"wire/internal/service/service_user"
)

var Provider = wire.NewSet(
	controller_user.NewUserController,
	wire.Bind(new(service_user.IUserService), new(*service_user.UserService)),
	wire.Bind(new(repo_user.IUserRepo), new(*repo_user.UserRepo)),
	wire.Struct(new(service_user.UserService), "*"),
	wire.Struct(new(repo_user.UserRepo), "*"),
)
