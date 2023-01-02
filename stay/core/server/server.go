package server

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"stay/core/app"
)

type Server struct {
	e        *echo.Echo
	config   *app.Config
	Handlers []Handler
}

func New(conf *app.Config, handlers []Handler) *Server {
	e := echo.New()

	var endpoints []Endpoint
	for _, h := range handlers {
		es := h.GetEndpoints()
		endpoints = append(endpoints, es...)
	}

	// TODO: this shows the Swagger documentation
	// modify the server configuration to allow
	// or disable it by configuration
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// TODO: this shows the register endpoints
	// 	modify to allow or disable by config
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	fmt.Println(string(data))

	// TODO: modify this to receive middlewares as parameters
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	for _, endpoint := range endpoints {
		e.Add(
			endpoint.GetMethod(),
			endpoint.GetPath(),
			endpoint.GetHandlerFunc())
	}

	return &Server{
		e:        e,
		config:   conf,
		Handlers: handlers,
	}
}

func (s *Server) Run() {
	s.e.Start(fmt.Sprintf(":%d", s.config.Port))
}
