package controller

import (
	"github.com/google/wire"
	"wire/internal/repo"
	"wire/internal/service"
)

var Provider = wire.NewSet(
	NewUserController,
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
	wire.Struct(new(service.UserService), "*"),
	wire.Struct(new(repo.UserRepo), "*"),
)
