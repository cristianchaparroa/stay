package server

import "github.com/labstack/echo/v4"

type Endpoint interface {
	GetPath() string
	GetMethod() string
	GetHandlerFunc() echo.HandlerFunc
}

type endpoint struct {
	path   string
	method string
	f      echo.HandlerFunc
}

func (e *endpoint) GetPath() string {
	return e.path
}

func (e *endpoint) GetMethod() string {
	return e.method
}

func (e *endpoint) GetHandlerFunc() echo.HandlerFunc {
	return e.f
}

func NewEndpoint(f echo.HandlerFunc, method, path string) Endpoint {
	return &endpoint{
		f:      f,
		path:   path,
		method: method,
	}
}

type Handler interface {
	GetEndpoints() []Endpoint
}
