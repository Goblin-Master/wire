//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire/internal/conf"
	"wire/internal/controller"
)

//go:generate wire
func WireApp() (*conf.HttpServer, error) {
	panic(wire.Build(
		wire.Struct(new(conf.HttpServer), "*"),
		conf.InitZap,
		conf.InitGin,
		conf.InitDB,
		controller.Provider,
	))
}
