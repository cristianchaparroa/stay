package server

import (
	"fmt"
	"stay/internal/app"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e      *echo.Echo
	config *app.Config
}

func New(conf *app.Config) *Server {
	e := echo.New()
	return &Server{e: e, config: conf}
}

func (s *Server) Run() {
	s.e.Start(fmt.Sprintf(":%d", s.config.Port))
}
