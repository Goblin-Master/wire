package main

import "wire/internal/conf"

func main() {
	conf.LoadConfig()
	app, err := WireApp()
	if err != nil {
		panic(err)
	}
	app.RunServer()
}
