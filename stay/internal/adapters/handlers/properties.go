package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stay/core"
	"stay/core/server"
	"stay/internal/ports"
)

type Properties struct {
	useCase ports.PropertyUseCase
}

func NewProperties(c ports.PropertyUseCase) server.Handler {
	return &Properties{useCase: c}
}

func (h *Properties) Create(c echo.Context) error {
	req, errReq := NewPropertyCreateRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	p, err := h.useCase.Save(req.ToDomain())
	if err != nil {
		e := core.NewError(err.Error(), http.StatusUnprocessableEntity)
		return c.JSON(e.Status, e)
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *Properties) GetAll(c echo.Context) error {
	req, errReq := NewPropertyGetAllRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	ps, err := h.useCase.GetAll(req.UID)

	if err != nil {
		e := core.NewError(err.Error(), http.StatusUnprocessableEntity)
		return c.JSON(e.Status, e)
	}

	response := NewPropertyListResponse(ps)
	return c.JSON(http.StatusOK, response)
}

func (h *Properties) GetByID(c echo.Context) error {
	req, errReq := NewPropertyGetByIDRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	p, err := h.useCase.GetProperty(req.UID, req.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	return c.JSON(http.StatusOK, NewPropertyResponse(p))
}

func (h *Properties) GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		server.NewEndpoint(h.GetAll, http.MethodGet, "/users/:uid/properties"),
		server.NewEndpoint(h.GetByID, http.MethodGet, "/users/:uid/properties/:id"),
		server.NewEndpoint(h.Create, http.MethodPost, "/users/properties"),
	}
}
