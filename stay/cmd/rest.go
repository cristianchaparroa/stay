package main

import (
	"stay/internal/adapters/server"
	"stay/internal/app"
)

func main() {
	c := app.NewAppConfig(8080)
	s := server.New(c)
	s.Run()
}
